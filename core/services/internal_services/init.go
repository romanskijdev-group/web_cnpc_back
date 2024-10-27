package servicesinternal

import (
	gormmodule "cnpc_backend/core/services/internal_services/gorm"
	pgxpool2 "cnpc_backend/core/services/internal_services/pgxpool"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)

func InitRedisModule(redisConfigObj *redismodule.RedisConfig) *redismodule.ModuleRedis {
	return redismodule.GetRedisModuleInstance(redisConfigObj)
}

func InitPgxPoolModule(configObj *pgxpool2.ConfigConnectPgxPool) *pgxpool.Pool {
	return pgxpool2.ConnectDB(configObj)
}

func InitGormService(configObj *gormmodule.ConfigConnectGorm) (*gorm.DB, *sql.DB) {
	return gormmodule.GormDatabaseConnect(configObj)
}
