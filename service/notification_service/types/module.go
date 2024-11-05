package types

import (
	usersdb "cnpc_backend/core/module/user/users/db"
	alicloudossapi "cnpc_backend/core/services/external_services/alicloud_oss_api"
	firebasepush "cnpc_backend/core/services/external_services/firebase_push"
	redismodule "cnpc_backend/core/services/internal_services/redis"
	"cnpc_backend/core/typescore"
	"html/template"
	interfacenotification "notification_service/notification/interface"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Modules struct {
	Notification interfacenotification.NotificationI
	BundleI18n   *i18n.Bundle
}

type DatabaseModuleI struct {
	//UserDevicePush pushuserdevicesdb.UserDevicePushDBI
	UsersActions usersdb.UsersProviderControlsDBI
	//SystemNotify   alertusernotifydb.AlertNotifyDBI
}

type Clients struct {
}

type TemplatesMailSystem struct {
	TfaEmailTemplate      *template.Template
	MailLoginPassTemplate *template.Template
	NewDeviceInfoTemplate *template.Template
	CustomNotifyTemplate  *template.Template
}

type InternalProviderControl struct {
	TemplatesMail  *TemplatesMailSystem
	Config         *typescore.Config
	FirebaseClient *firebasepush.FirebaseClient

	DatabasePull *pgxpool.Pool
	RedisClient  *redismodule.ModuleRedis

	Modules  Modules
	Database DatabaseModuleI
	Clients  Clients
	Storage  alicloudossapi.AliCloudOSSStorageI
}
