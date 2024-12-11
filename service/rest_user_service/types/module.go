package types

import (
	ipdetectorinterface "cnpc_backend/core/common/ip_detector/interface"
	restauthcore "cnpc_backend/core/module/rest_auth"
	usersdb "cnpc_backend/core/module/user/users/db"
	protoobj "cnpc_backend/core/proto"
	awss3api "cnpc_backend/core/services/external_services/aws_s3_api"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Clients struct {
	UserAccountServiceProto  protoobj.UserAccountServiceProtoClient
	NotificationServiceProto protoobj.NotificationServiceProtoClient
}

type Modules struct {
	RestAuth         *restauthcore.ModuleRestAuth
	IPDetectorModule ipdetectorinterface.IPdetectorI
}

type DatabaseI struct {
	UsersDB usersdb.UsersProviderControlsDBI
}

type InternalProviderControl struct {
	Config       *typescore.Config
	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis
	Modules      Modules
	Clients      Clients
	Storage      awss3api.AWSS3CloudStorageI
	Database     DatabaseI
}
