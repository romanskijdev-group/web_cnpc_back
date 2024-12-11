package marshalleruseralerts

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func UsersAlertsDeserialization(obj *protoobj.UserAlertMsg) *typescore.UserSystemAlerts {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.UserSystemAlerts{
		SystemID:   d.OptionalString(obj.SystemID),
		SerialID:   d.OptionalUint64(obj.SerialID),
		CreatedAt:  d.OptionalTime(obj.CreatedAt),
		UserID:     d.OptionalString(obj.UserID),
		Reading:    d.OptionalBool(obj.Reading),
		NotifyType: (*typescore.NotifyCategory)(d.OptionalString(obj.NotifyType)),
		Title:      d.OptionalString(obj.Title),
		Message:    d.OptionalString(obj.Message),
		Link:       d.OptionalString(obj.Link),
		DeepLinkID: d.OptionalUint64(obj.DeepLinkID),
	}
}

func UsersAlertsMsgListDeserialization(obj *protoobj.UserAlertMsgList) []*typescore.UserSystemAlerts {
	if obj == nil {
		return nil
	}
	var res []*typescore.UserSystemAlerts
	for _, item := range obj.UsersAlerts {
		deserializedItem := UsersAlertsDeserialization(item)
		if deserializedItem != nil {
			res = append(res, deserializedItem)
		}
	}
	return res
}
