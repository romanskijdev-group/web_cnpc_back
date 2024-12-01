package types

import (
	alertsdb "cnpc_backend/core/module/notification/user_alerts/db"
	restauthcore "cnpc_backend/core/module/rest_auth"
	usersdb "cnpc_backend/core/module/user/users/db"
	protoobj "cnpc_backend/core/proto"
	awss3api "cnpc_backend/core/services/external_services/aws_s3_api"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Clients struct {
	NotificationServiceProto protoobj.NotificationServiceProtoClient
}

type Modules struct {
	RestAuth   *restauthcore.ModuleRestAuth
	BundleI18n *i18n.Bundle
}

type DatabaseModuleI struct {
	UsersActions usersdb.UsersProviderControlsDBI
	UserAlerts   alertsdb.UsersAlertsDBI
	//ReferralBonus      referralbonusesdb.ReferralBonusesDBI
	//UsersSubscriptions userssubsdb.UsersSubscriptionDBI
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
