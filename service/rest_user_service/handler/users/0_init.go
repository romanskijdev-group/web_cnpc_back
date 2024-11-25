package usershandler

import (
	usershandlerinterface "cnpc_backend/rest_user_service/handler/users/interface"
	"cnpc_backend/rest_user_service/types"
)

type HandlerAccount struct {
	ipc         *types.InternalProviderControl
	UserProfile usershandlerinterface.UsersI
}

func NewHandlerUsers(ipc *types.InternalProviderControl) *HandlerAccount {
	return &HandlerAccount{
		ipc:         ipc,
		UserProfile: newHandlerAccount(ipc, "UserProfile").(usershandlerinterface.UsersI),
	}
}

// Обобщённая функция инициализации
func newHandlerAccount(ipc *types.InternalProviderControl, accountType string) interface{} {
	switch accountType {
	case "UserProfile":
		return &HandlerUsers{ipc: ipc}
	default:
		return nil
	}
}
