package user

import (
	usersdb "cnpc_backend/core/module/user/users/db"
	"cnpc_backend/core/typescore"
)

func InitUsersModuleDB(configModule *typescore.ModuleDBConfig) usersdb.UsersProviderControlsDBI {
	return usersdb.NewUsersDB(configModule)
}
