package notification

import (
	interfacenotification "notification_service/notification/interface"
	"notification_service/types"
)

type ModuleNotification struct {
	ipc *types.InternalProviderControl
}

func NewModuleNotification(ipc *types.InternalProviderControl) interfacenotification.NotificationI {
	return &ModuleNotification{
		ipc: ipc,
	}
}
