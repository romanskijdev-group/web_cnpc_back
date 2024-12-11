package usersdb

import (
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

var TableName = "users"

type ModuleDB struct {
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
}

func NewUsersDB(configModule *typescore.ModuleDBConfig) UsersProviderControlsDBI {
	return &ModuleDB{
		DatabasePull: configModule.DatabasePull,
		RedisClient:  configModule.RedisClient,
	}
}

type UsersProviderControlsDBI interface {
	UpdateUserLastLoginInfoDB(ctx context.Context, UsersProviderControlObj *typescore.UsersProviderControl) *typescore.WEvent
	DeleteUserDB(ctx context.Context, UsersProviderControlParams *typescore.UsersProviderControl) *typescore.WEvent
	UpdateUserDB(ctx context.Context, paramsUpdate *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent)
	CreateUserDB(ctx context.Context, UsersProviderControlObj *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent)
	GetUsersListDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl, likeFields map[string]string, offset *uint64, limit *uint64) ([]*typescore.UsersProviderControl, *typescore.WEvent)
	GetUserDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl) (*typescore.UsersProviderControl, *typescore.WEvent)
	GetUsersCountDB(ctx context.Context, paramsFiltering *typescore.UsersProviderControl, likeFields map[string]string) (uint64, *typescore.WEvent)
	GetUsersStatisticsByDateDB(ctx context.Context, paramsFiltering *typescore.TimePeriod, statType *typescore.UserStatisticsType) ([]*typescore.CountByDateStatisticsResponse, *typescore.WEvent)
	UpdateUserBalanceDB(tx pgx.Tx, ctx context.Context, obj *typescore.UsersProviderControl, amount *decimal.Decimal) (pgx.Tx, error)
	UpdateUserAvatarURLDB(ctx context.Context, userSystemID *string, avatarURL string) *typescore.WEvent
}
