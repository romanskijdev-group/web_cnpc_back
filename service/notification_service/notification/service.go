package notification

import (
	marshallernotification "cnpc_backend/core/module/notification/marshaller"
	protoobj "cnpc_backend/core/proto"
	"context"
	"log"
	interfacenotification "notification_service/notification/interface"
	"notification_service/types"
)

func NewNotificationServiceProto(ipc *types.InternalProviderControl) protoobj.NotificationServiceProtoServer {
	return &NotificationServiceProto{
		ipc:                ipc,
		ModuleNotification: NewModuleNotification(ipc),
	}
}

type NotificationServiceProto struct {
	ipc                *types.InternalProviderControl
	ModuleNotification interfacenotification.NotificationI
	protoobj.UnimplementedNotificationServiceProtoServer
}

func (t *NotificationServiceProto) NotifyUser(ctx context.Context, prObj *protoobj.NotifyParams) (*protoobj.Empty, error) {
	notifyParamsObj := marshallernotification.NotifyParamsDeserialization(prObj)

	err := t.ModuleNotification.NotifyRouting(notifyParamsObj)
	if err != nil {
		log.Println("🔴 error NotifyUser NotifyRouting: ", err)
		return nil, err
	}
	return nil, nil
}
