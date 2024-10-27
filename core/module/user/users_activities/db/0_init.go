package userslogindb

import (
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var TableName = "login_activities"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewUsersLoginDB(configModule *typescore.ModuleDBConfig) UsersLoginDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type UsersLoginDBI interface {
	// GetUsersStatisticsByDateDB получение статистик пользователей по датам
	GetUsersStatisticsByDateDB(ctx context.Context, paramsFiltering *typescore.TimePeriod) ([]*typescore.CountByDateStatisticsResponse, *typescore.WEvent)

	// AddUsersLoginStatisticsDB Добавление записи количества активных пользователей
	AddUsersLoginStatisticsDB(ctx context.Context, paramsFiltering *typescore.TimePeriod) *typescore.WEvent
}
