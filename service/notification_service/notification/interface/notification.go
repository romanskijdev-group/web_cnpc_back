package interfacenotification

import "cnpc_backend/core/typescore"

type NotificationI interface {
	// Routing уведомлений
	NotifyRouting(notifyParams *typescore.NotifyParams) error
}
