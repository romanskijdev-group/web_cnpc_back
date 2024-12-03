package notifications

import (
	notificationsinterface "cnpc_backend/rest_user_service/handler/notifications/interface"
	"cnpc_backend/rest_user_service/types"
)

type HandlerNotifications struct {
	ipc        *types.InternalProviderControl
	UserAlerts notificationsinterface.UserAlertsI
}

func NewHandlerNotifications(ipc *types.InternalProviderControl) *HandlerNotifications {
	return &HandlerNotifications{
		ipc:        ipc,
		UserAlerts: newHandlerAccount(ipc, "UserAlerts").(notificationsinterface.UserAlertsI),
	}
}

// Обобщённая функция инициализации
func newHandlerAccount(ipc *types.InternalProviderControl, accountType string) interface{} {
	switch accountType {
	case "UserAlerts":
		return &HandlerUserAlerts{ipc: ipc}
	default:
		return nil
	}
}
