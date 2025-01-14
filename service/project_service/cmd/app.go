package main

import (
	"cnpc_backend/core/config"
	grpccore "cnpc_backend/core/grpc_core/grpc"
	alertsdb "cnpc_backend/core/module/notification/user_alerts/db"
	restauthcore "cnpc_backend/core/module/rest_auth"
	usersdb "cnpc_backend/core/module/user/users/db"
	protoobj "cnpc_backend/core/proto"
	awss3api "cnpc_backend/core/services/external_services/aws_s3_api"
	internalservices "cnpc_backend/core/services/internal_services"
	"cnpc_backend/core/services/internal_services/pgxpool"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	grpcclients "projectservice/grpc_clients"
	"projectservice/types"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}

	serverBind := cfg.Server.UserService.Bind
	serverPort := cfg.Server.UserService.Port
	if serverBind == "" || serverPort == 0 {
		log.Fatal(errors.New("projects_service: missed bind or port params"))
	}
	serverStartURI := fmt.Sprintf("%s:%d",
		serverBind,
		serverPort)

	log.Println("⭐️⭐️⭐️ projects_service started: ", serverStartURI)

	ipc := initInternalProvider(cfg)

	protoOpt := grpcclients.CreateDialOptionsProto()
	ipc.Clients = types.Clients{
		NotificationServiceProto: grpcclients.InitClientNotificationServiceProto(protoOpt, cfg),
	}

	ipc = initModules(cfg, ipc)

	startGRPC(ipc, serverStartURI)
}

func i8nInit() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./locale/locale.en.toml")
	bundle.MustLoadMessageFile("./locale/locale.ru.toml")
	return bundle
}

func startGRPC(ipc *types.InternalProviderControl, addressService string) {
	server, err := grpccore.CreateServerGRPC(nil, nil)
	if err != nil {
		log.Fatalf("🔴 Failed to create gRPC server: %v", err)
	}
	// Регистрация сервисов на сервере
	protoobj.RegisterUserAccountServiceProtoServer(server, useraccount.NewUserAccountServiceProto(ipc))

	// Создание и регистрация сервера проверки состояния
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)

	lis, err := net.Listen("tcp", addressService)
	if err != nil {
		log.Fatalf("🔴 Failed to listen: %v", err)
	}
	log.Printf("🟢 Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("🔴 Failed to serve: %v", err)
	}
}

func initModules(config *typescore.Config, ipc *types.InternalProviderControl) *types.InternalProviderControl {
	configModules := &typescore.ModuleDBConfig{
		DatabasePull: ipc.DatabasePull,
		RedisClient:  ipc.RedisClient,
		ConfigGlobal: config,
	}

	ipc.Database = types.DatabaseModuleI{
		UsersActions: usersdb.NewUsersDB(configModules),
		UserAlerts:   alertsdb.NewUserAlertsDB(configModules),
		//ReferralBonus:      referralbonusesdb.NewReferralBonusesDB(configModules),
		//UsersSubscriptions: userssubsdb.NewUsersSubscriptionDB(configModules),
	}
	ipc.Modules = types.Modules{
		RestAuth: restauthcore.InitNewModule(configModules),
	}
	ipc.Modules.BundleI18n = i8nInit()
	return ipc
}

func initInternalProvider(config *typescore.Config) *types.InternalProviderControl {
	redisClient := internalservices.InitRedisModule(&redismodule.RedisConfig{
		Host:     config.Redis.Addr,
		Port:     config.Redis.Port,
		User:     config.Redis.Username,
		Password: config.Redis.Password,
	})
	databasePg := internalservices.InitPgxPoolModule(&pgxpool.ConfigConnectPgxPool{
		Host:     config.Storage.Host,
		Port:     config.Storage.Port,
		User:     config.Storage.Username,
		Password: config.Storage.Password,
		Name:     config.Storage.DBName,
	})

	return &types.InternalProviderControl{
		Config:       config,
		DatabasePull: databasePg,
		RedisClient:  redisClient,
		Storage: awss3api.NewAWSS3CloudStorage(awss3api.StorageConfigSt{
			Key:      config.CloudStorage.Key,
			Secret:   config.CloudStorage.Secret,
			Region:   config.CloudStorage.Region,
			Bucket:   config.CloudStorage.Bucket,
			Endpoint: config.CloudStorage.Endpoint,
			RunMode:  config.CloudStorage.RunMode,
		}),
	}
}
