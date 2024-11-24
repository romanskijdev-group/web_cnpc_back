package main

import (
	"cnpc_backend/core/config"
	restauthcore "cnpc_backend/core/module/rest_auth"
	usersdb "cnpc_backend/core/module/user/users/db"
	awss3api "cnpc_backend/core/services/external_services/aws_s3_api"
	servicesinternal "cnpc_backend/core/services/internal_services"
	"cnpc_backend/core/services/internal_services/pgxpool"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	grpcclients "cnpc_backend/rest_user_service/grpc_clients"
	authuser "cnpc_backend/rest_user_service/handler/auth"
	userprofile "cnpc_backend/rest_user_service/handler/profile"
	usershandler "cnpc_backend/rest_user_service/handler/users"
	"cnpc_backend/rest_user_service/types"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("⭐️⭐️⭐️ starting rest_user_service...")
	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}

	// ** Инициализация модулей
	ipc := initInternalProvider(cfg)
	protoOpt := grpcclients.CreateDialOptionsProto()
	ipc.Clients = types.Clients{
		UserAccountServiceProto:  grpcclients.InitClientUserAccountServiceProto(protoOpt, cfg),
		NotificationServiceProto: grpcclients.InitClientNotificationServiceProto(protoOpt, cfg),
	}
	ipc = initModules(cfg, ipc)
	startRest(ipc)
}

func initModules(configObj *typescore.Config, ipc *types.InternalProviderControl) *types.InternalProviderControl {
	configModules := &typescore.ModuleDBConfig{
		DatabasePull: ipc.DatabasePull,
		RedisClient:  ipc.RedisClient,
		ConfigGlobal: configObj,
	}

	ipc.Modules = types.Modules{
		RestAuth: restauthcore.InitNewModule(configModules),
	}

	ipc.Database = types.DatabaseI{
		UsersDB: usersdb.NewUsersDB(configModules),
	}

	return ipc
}

func initInternalProvider(config *typescore.Config) *types.InternalProviderControl {
	redisClient := servicesinternal.InitRedisModule(&redismodule.RedisConfig{
		Host:     config.Redis.Addr,
		Port:     config.Redis.Port,
		User:     config.Redis.Username,
		Password: config.Redis.Password,
	})
	databasePg := servicesinternal.InitPgxPoolModule(&pgxpool.ConfigConnectPgxPool{
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

// Регистрация обработчиков
func registerRouters(ipc *types.InternalProviderControl, router *chi.Mux) {
	// Регистрация обработчика авторизации
	authUser := authuser.NewAuthUser(ipc)
	authUser.RegisterAuthByToken(router)

	profile := userprofile.NewHandlerAccount(ipc)
	profile.RegisterProfile(router)

	users := usershandler.NewHandlerUsers(ipc)
	users.RegisterUsers(router)
}

// запуск сервера REST API
func startRest(ipc *types.InternalProviderControl) {
	router := chi.NewRouter()
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Список разрешенных origin
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Cache-Control", "X-Requested-With", "x-api"},
		ExposedHeaders:   []string{"Link", "Cache-Control"},
		AllowCredentials: false,
		MaxAge:           300, // Максимальное время жизни предварительных запросов в секундах
	})

	router.Use(corsOptions.Handler)
	router.Use(middleware.RealIP)

	// Регистрация обработчиков
	registerRouters(ipc, router)

	serverBind := ipc.Config.Server.RESTUserService.Bind
	serverPort := ipc.Config.Server.RESTUserService.Port
	if serverBind == "" || serverPort == 0 {
		log.Fatal(errors.New("rest_user_service: missed bind or port params"))
	}
	serverStartURI := fmt.Sprintf("%s:%d",
		serverBind,
		serverPort)

	listener, err := net.Listen("tcp", serverStartURI)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Minute,
		WriteTimeout: 15 * time.Minute,
		IdleTimeout:  15 * time.Minute,
	}

	log.Println("✅ Application initialized and started: ", serverStartURI)

	if err := httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("Server shutdown")
		default:
			log.Fatal(err)
		}
	}
}
