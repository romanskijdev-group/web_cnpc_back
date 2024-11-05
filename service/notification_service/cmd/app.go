package main

import (
	"cnpc_backend/core/config"
	grpccore "cnpc_backend/core/grpc_core/grpc"
	usersdb "cnpc_backend/core/module/user/users/db"
	protoobj "cnpc_backend/core/proto"
	internalservices "cnpc_backend/core/services/internal_services"
	"cnpc_backend/core/services/internal_services/pgxpool"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"notification_service/loader"
	"notification_service/notification"
	"notification_service/types"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}

	addressService := fmt.Sprintf("%s:%d", cfg.Server.NotificationsService.Bind, cfg.Server.NotificationsService.Port)
	log.Println("⚡️🚀⚡️ notifications_service started: ", addressService)

	ipc := initInternalProvider(cfg)

	ipc = initModules(cfg, ipc)
	startGrps(ipc, addressService)
}

func i8nInit() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./locale/locale.en.toml")
	bundle.MustLoadMessageFile("./locale/locale.ru.toml")
	return bundle
}

func startGrps(ipc *types.InternalProviderControl, addressService string) {
	server, err := grpccore.CreateServerGRPC(nil, nil)
	if err != nil {
		log.Fatalf("🔴 Failed to create gRPC server: %v", err)
	}

	// Регистрация сервисов на сервере
	protoobj.RegisterNotificationServiceProtoServer(server, notification.NewNotificationServiceProto(ipc))

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
		//UserDevicePush: pushuserdevicesdb.NewUserDevicePushDB(configModules),
		UsersActions: usersdb.NewUsersDB(configModules),
		//SystemNotify:   alertusernotifydb.NewAlertNotifyDB(configModules),
	}
	ipc.Modules.Notification = notification.NewModuleNotification(ipc)
	ipc.Modules.BundleI18n = i8nInit()
	ipc.TemplatesMail = loader.LoadMailTemplates()
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

	//firebaseClient := externalservices.InitFirebaseService(&firebasepush.FirebaseConfig{
	//	FcmServerToken: &config.Firebase.CredentialsServerToken,
	//})

	//storageModule := alicloudossapi.NewAliCloudOSSStorage(alicloudossapi.StorageConfigSt{
	//	Key:      config.AliCloudOSSStorage.Key,
	//	Secret:   config.AliCloudOSSStorage.Secret,
	//	Endpoint: config.AliCloudOSSStorage.Endpoint,
	//	Bucket:   config.AliCloudOSSStorage.Bucket,
	//})

	return &types.InternalProviderControl{
		Config:       config,
		DatabasePull: databasePg,
		RedisClient:  redisClient,
		//FirebaseClient: firebaseClient,
		//Storage: storageModule,
	}
}
