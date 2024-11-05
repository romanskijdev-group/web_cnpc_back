package marshallernotification

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func NotifyParamsSerialization(obj *typescore.NotifyParams) *protoobj.NotifyParams {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	systemUserIDs := make([]*wrapperspb.StringValue, 0)
	for _, v := range obj.SystemUserIDs {
		systemUserIDs = append(systemUserIDs, s.StringToWrapperStringValue(v))
	}

	additionsObject := NotifyAdditionsObjectSerialization(obj.AdditionsObject)

	return &protoobj.NotifyParams{
		Text:            s.StringToWrapperStringValue(obj.Text),
		Title:           s.StringToWrapperStringValue(obj.Title),
		SystemUserIDs:   systemUserIDs,
		MailAddress:     s.StringToWrapperStringValue(obj.MailAddress),
		Category:        mapNotifyCategoryTypeSerialization(obj.Category),
		IsEmail:         obj.IsEmail,
		Emergency:       obj.Emergency,
		AdditionsObject: additionsObject,
	}
}

func NotifyAdditionsObjectSerialization(obj *typescore.NotifyAdditionsObject) *protoobj.NotifyAdditionsObject {
	if obj == nil {
		return nil
	}
	//bearerCheques := marshallerpaymentcheckpayablebearer.BearerChequeSerialization(obj.BearerChequeObj)
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.NotifyAdditionsObject{
		//BearerChequeObj:     bearerCheques,
		SubmittedBy: s.StringToWrapperStringValue(obj.SubmittedBy),
	}
}

func mapNotifyCategoryTypeSerialization(typeNotifyCategory *typescore.NotifyCategory) protoobj.NotifyCategory {
	if typeNotifyCategory == nil {
		return protoobj.NotifyCategory_NotifyCategory_NULL
	}
	switch *typeNotifyCategory {
	case typescore.BearerChequeNotifyCategory:
		return protoobj.NotifyCategory_NotifyCategory_bearer_cheque
	case typescore.ChatsMessageNotifyCategory:
		return protoobj.NotifyCategory_NotifyCategory_chats
	case typescore.InfoNotifyCategory:
		return protoobj.NotifyCategory_NotifyCategory_info
	case typescore.TemporaryPasswordNotifyCategory:
		return protoobj.NotifyCategory_NotifyCategory_temporary_password
	case typescore.DeviceNewNotifyCategory:
		return protoobj.NotifyCategory_NotifyCategory_device_new
	default:
		return protoobj.NotifyCategory_NotifyCategory_NULL
	}
}
