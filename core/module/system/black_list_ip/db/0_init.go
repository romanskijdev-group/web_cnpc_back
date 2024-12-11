package blacklistipdb

import (
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TableName = "blacklist_ip"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewBlackListIPDB(configModule *typescore.ModuleDBConfig) BlackListIPDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type BlackListIPDBI interface { // Получение списка Ip в черном списке
	GetBlackListIPDB(ctx context.Context, paramsFiltering *typescore.BlackListIP) ([]*typescore.BlackListIP, *typescore.WEvent)
	// Добавление IP в черный список
	AddIPToBlackListDB(ctx context.Context, blackIPItem *typescore.BlackListIP) (*typescore.BlackListIP, *typescore.WEvent)
}
