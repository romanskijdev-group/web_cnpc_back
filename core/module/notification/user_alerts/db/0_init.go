package alertsdb

import (
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var TableName = "user_alerts"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewUserAlertsDB(configModule *typescore.ModuleDBConfig) UsersAlertsDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type UsersAlertsDBI interface {
	GetUserAlertsListDB(ctx context.Context, paramsFiltering *typescore.UserSystemAlerts, likeFields map[string]string, offset *uint64, limit *uint64) ([]*typescore.UserSystemAlerts, *typescore.WEvent)
	CreateUserAlertDB(ctx context.Context, paramsCreate *typescore.UserSystemAlerts) (*typescore.UserSystemAlerts, *typescore.WEvent)
	UpdateUserAlertDB(ctx context.Context, paramsUpdate *typescore.UserSystemAlerts) ([]*typescore.UserSystemAlerts, *typescore.WEvent)
	DeleteUserAlertDB(ctx context.Context, paramsDelete *typescore.UserSystemAlerts) *typescore.WEvent
	GetUserAlertsCountDB(ctx context.Context, paramsFiltering *typescore.UserSystemAlerts, likeFields map[string]string) (uint64, *typescore.WEvent)
}
