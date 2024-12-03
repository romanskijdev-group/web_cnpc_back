package marshalleruseralerts

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func UserAlertSerialization(obj *typescore.UserSystemAlerts) *protoobj.UserAlertMsg {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UserAlertMsg{
		SystemID:   s.StringToWrapperStringValue(obj.SystemID),
		SerialID:   s.Uint64ToWrapperUInt64Value(obj.SerialID),
		CreatedAt:  s.TimePtrToTimestampPB(obj.CreatedAt),
		Reading:    s.BoolToWrapperBoolValue(obj.Reading),
		NotifyType: s.StringToWrapperStringValue((*string)(obj.NotifyType)),
		Title:      s.StringToWrapperStringValue(obj.Title),
		Message:    s.StringToWrapperStringValue(obj.Message),
		Link:       s.StringToWrapperStringValue(obj.Link),
		DeepLinkID: s.Uint64ToWrapperUInt64Value(obj.DeepLinkID),
	}
}

func UserAlertMsgListSerialization(obj []*typescore.UserSystemAlerts) *protoobj.UserAlertMsgList {
	if obj == nil {
		return nil
	}
	var res []*protoobj.UserAlertMsg
	for _, item := range obj {
		protoItem := UserAlertSerialization(item)
		if protoItem != nil {
			res = append(res, protoItem)
		}
	}
	return &protoobj.UserAlertMsgList{
		UsersAlerts: res,
	}
}
