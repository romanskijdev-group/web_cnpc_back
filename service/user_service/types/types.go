package types

import (
	restauthcore "cnpc_backend/core/module/rest_auth"
	userssubsdb "cnpc_backend/core/module/user/subscription/db"
	usersdb "cnpc_backend/core/module/user/users/db"
	awss3api "cnpc_backend/core/services/external_services/aws_s3_api"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Clients struct {
}

type Modules struct {
	RestAuth *restauthcore.ModuleRestAuth
}

type DatabaseModuleI struct {
	UsersActions usersdb.UsersProviderControlsDBI
	//ReferralBonus      referralbonusesdb.ReferralBonusesDBI
	UsersSubscriptions userssubsdb.UsersSubscriptionDBI
}

type InternalProviderControl struct {
	Config *typescore.Config

	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
	Storage      awss3api.AWSS3CloudStorageI
	Clients      Clients
	Modules      Modules
	Database     DatabaseModuleI
}
