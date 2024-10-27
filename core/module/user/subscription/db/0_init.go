package userssubsdb

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"zod_backend_dev/core/models"
	redismodule "zod_backend_dev/core/services/internal_services/redis"
)

var TableName = "users_subscriptions"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewUsersSubscriptionDB(configModule *models.ModuleDBConfig) UsersSubscriptionDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type UsersSubscriptionDBI interface {
	UpdateUserSubscriptionDB(ctx context.Context, paramsUpdate *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent)
	CreateUserSubscriptionDB(ctx context.Context, userObj *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent)
	GetUsersSubscriptionsListDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions, likeFields map[string]string, offset *uint64, limit *uint64) ([]*models.UsersSubscriptions, *models.WEvent)
	GetUserSubscriptionDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions) (*models.UsersSubscriptions, *models.WEvent)
	GetUsersLimitsCountDB(ctx context.Context, paramsFiltering *models.UsersSubscriptions, likeFields map[string]string) (uint64, *models.WEvent)
}
