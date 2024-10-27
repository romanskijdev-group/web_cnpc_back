package restauthcore

import (
	restauthcoreinterface "cnpc_backend/core/module/rest_auth/interface"
	"cnpc_backend/core/module/user"
	"cnpc_backend/core/typescore"
)

type ModuleRestAuth struct {
	ConfigModule    *typescore.ModuleDBConfig
	CheckerRestAuth restauthcoreinterface.CheckerRestAuthI
	TokenRestAuth   restauthcoreinterface.TokenRestAuthI
}

func InitNewModule(configModule *typescore.ModuleDBConfig) *ModuleRestAuth {
	userDb := user.InitUsersModuleDB(configModule)

	tokenRestAutModule := &TokenRestAuth{
		ConfigG: configModule.ConfigGlobal,
	}

	checkerRestAuthModule := &CheckerRestAuth{
		ConfigG: configModule.ConfigGlobal,
		UsersDB: userDb,
	}

	return &ModuleRestAuth{
		CheckerRestAuth: checkerRestAuthModule,
		TokenRestAuth:   tokenRestAutModule,
		ConfigModule:    configModule,
	}
}
