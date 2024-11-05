package authuser

import (
	interfaceauth "cnpc_backend/rest_user_service/handler/auth/interface"
	"cnpc_backend/rest_user_service/types"
)

type AuthUser struct {
	ipc           *types.InternalProviderControl
	AuthUserToken interfaceauth.AuthUserI
}

func NewAuthUser(ipc *types.InternalProviderControl) *AuthUser {
	return &AuthUser{
		ipc:           ipc,
		AuthUserToken: newAuthUser(ipc, "AuthUserToken").(interfaceauth.AuthUserI),
	}
}

// Обобщённая функция инициализации
func newAuthUser(ipc *types.InternalProviderControl, accountType string) interface{} {
	switch accountType {
	case "AuthUserToken":
		return &HandlerAuthByToken{ipc: ipc}
	default:
		return nil
	}
}
