package typescore

import (
	"cnpc_backend/core/services/internal_services/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ModuleDBConfig struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
	ConfigGlobal *Config
}
