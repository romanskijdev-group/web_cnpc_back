package userssubsdb

import (
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var TableName = "users_subscriptions"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewUsersSubscriptionDB(configModule *typescore.ModuleDBConfig) UsersSubscriptionDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type UsersSubscriptionDBI interface {
	UpdateUserSubscriptionDB(ctx context.Context, paramsUpdate *typescore.UsersSubscriptions) (*typescore.UsersSubscriptions, *typescore.WEvent)
	CreateUserSubscriptionDB(ctx context.Context, userObj *typescore.UsersSubscriptions) (*typescore.UsersSubscriptions, *typescore.WEvent)
	GetUsersSubscriptionsListDB(ctx context.Context, paramsFiltering *typescore.UsersSubscriptions, likeFields map[string]string, offset *uint64, limit *uint64) ([]*typescore.UsersSubscriptions, *typescore.WEvent)
	GetUserSubscriptionDB(ctx context.Context, paramsFiltering *typescore.UsersSubscriptions) (*typescore.UsersSubscriptions, *models.WEvent)
	GetUsersLimitsCountDB(ctx context.Context, paramsFiltering *typescore.UsersSubscriptions, likeFields map[string]string) (uint64, *models.WEvent)
}
