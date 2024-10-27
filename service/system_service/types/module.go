package types

import (
	usersdb "cnpc_backend/core/module/user/users/db"
	userslogindb "cnpc_backend/core/module/user/users_activities/db"
	exchangerateapi "cnpc_backend/core/services/external_services/exchangerate_api"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"gorm.io/gorm"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Clients struct {
}

type Modules struct {
	ExchangeRateAPI exchangerateapi.ExchangeRateI
}

type DatabaseModuleI struct {
	UsersActions  usersdb.UsersProviderControlsDBI
	UsersActivity userslogindb.UsersLoginDBI
}

type InternalProviderControl struct {
	Config *typescore.Config

	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
	GormDatabase *gorm.DB

	Clients  Clients
	Modules  Modules
	Database DatabaseModuleI
}
