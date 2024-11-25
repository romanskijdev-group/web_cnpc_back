package userprofile

import (
	interfaceuserprofile "cnpc_backend/rest_user_service/handler/profile/interface"
	"cnpc_backend/rest_user_service/types"
)

type HandlerAccount struct {
	ipc         *types.InternalProviderControl
	UserProfile interfaceuserprofile.UserProfileI
}

func NewHandlerAccount(ipc *types.InternalProviderControl) *HandlerAccount {
	return &HandlerAccount{
		ipc:         ipc,
		UserProfile: newHandlerAccount(ipc, "UserProfile").(interfaceuserprofile.UserProfileI),
	}
}

// Обобщённая функция инициализации
func newHandlerAccount(ipc *types.InternalProviderControl, accountType string) interface{} {
	switch accountType {
	case "UserProfile":
		return &HandlerUserProfile{ipc: ipc}
	default:
		return nil
	}
}
