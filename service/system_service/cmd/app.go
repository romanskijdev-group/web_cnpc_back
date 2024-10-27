package main

import (
	"cnpc_backend/core/config"
	"cnpc_backend/core/module/user"
	userslogindb "cnpc_backend/core/module/user/users_activities/db"
	externalservices "cnpc_backend/core/services/external_services"
	exchangerateapi "cnpc_backend/core/services/external_services/exchangerate_api"
	internalservices "cnpc_backend/core/services/internal_services"
	gormmodule "cnpc_backend/core/services/internal_services/gorm"
	"cnpc_backend/core/services/internal_services/pgxpool"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	backgroundjob "system_service/jobs"
	"system_service/types"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Print(fmt.Errorf("failed to load config: %w", err))
	}

	// ** Инициализация модулей
	ipc := initInternalProvider(cfg)
	ipc.Clients = types.Clients{}
	ipc = initModules(cfg, ipc)

	backgroundJobModule := backgroundjob.NewBackgroundJobModule(ipc)
	backgroundJobModule.StartJobsAll()

	log.Println("✅ Start background jobs service...")
	// Ожидание сигнала для корректного завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

func initModules(configObj *typescore.Config, ipc *types.InternalProviderControl) *types.InternalProviderControl {
	configModules := &typescore.ModuleDBConfig{
		DatabasePull: ipc.DatabasePull,
		RedisClient:  ipc.RedisClient,
		ConfigGlobal: configObj,
	}

	ipc.Database = types.DatabaseModuleI{
		UsersActions:  user.InitUsersModuleDB(configModules),
		UsersActivity: userslogindb.NewUsersLoginDB(configModules),
	}

	ipc.Modules.ExchangeRateAPI = externalservices.InitExchangeRateAPIModule(exchangerateapi.ExchangeRateConfigSt{
		ExchangeRateAPIKey: configObj.ExchangeRate.ExchangeRateAPIKey,
		ExchangeRateAPIURL: configObj.ExchangeRate.ExchangeRateAPIURL,
	})

	return ipc
}

func initInternalProvider(configObj *typescore.Config) *types.InternalProviderControl {
	redisClient := internalservices.InitRedisModule(&redismodule.RedisConfig{
		Host:     configObj.Redis.Addr,
		Port:     configObj.Redis.Port,
		User:     configObj.Redis.Username,
		Password: configObj.Redis.Password,
	})
	databasePg := internalservices.InitPgxPoolModule(&pgxpool.ConfigConnectPgxPool{
		Host:     configObj.Storage.Host,
		Port:     configObj.Storage.Port,
		User:     configObj.Storage.Username,
		Password: configObj.Storage.Password,
		Name:     configObj.Storage.DBName,
	})

	gormDataBase, _ := internalservices.InitGormService(&gormmodule.ConfigConnectGorm{
		Host:     configObj.Storage.Host,
		Port:     configObj.Storage.Port,
		User:     configObj.Storage.Username,
		Password: configObj.Storage.Password,
		Name:     configObj.Storage.DBName,
	})
	return &types.InternalProviderControl{
		Config:       configObj,
		DatabasePull: databasePg,
		RedisClient:  redisClient,
		GormDatabase: gormDataBase,
	}
}
