package marshallernotification

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func NotifyParamsDeserialization(obj *protoobj.NotifyParams) *typescore.NotifyParams {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	systemUserIDs := make([]*string, 0)
	for _, v := range obj.SystemUserIDs {
		systemUserIDs = append(systemUserIDs, d.OptionalString(v))
	}

	additionsObject := NotifyAdditionsObjectDeserialization(obj.AdditionsObject)
	return &typescore.NotifyParams{
		Text:            d.OptionalString(obj.Text),
		Title:           d.OptionalString(obj.Title),
		SystemUserIDs:   systemUserIDs,
		MailAddress:     d.OptionalString(obj.MailAddress),
		Category:        mapNotifyCategoryTypeDeserialization(obj.Category),
		IsEmail:         obj.IsEmail,
		Emergency:       obj.Emergency,
		AdditionsObject: additionsObject,
	}
}

func NotifyAdditionsObjectDeserialization(obj *protoobj.NotifyAdditionsObject) *typescore.NotifyAdditionsObject {
	if obj == nil {
		return nil
	}
	//bearerCheque := marshallerpaymentcheckpayablebearer.BearerChequeDeserialization(obj.BearerChequeObj)
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.NotifyAdditionsObject{
		//BearerChequeObj: bearerCheque,
		SubmittedBy: d.OptionalString(obj.SubmittedBy),
	}
}

func mapNotifyCategoryTypeDeserialization(typeNotifyCategory protoobj.NotifyCategory) *typescore.NotifyCategory {
	switch typeNotifyCategory {
	case protoobj.NotifyCategory_NotifyCategory_bearer_cheque:
		t := typescore.BearerChequeNotifyCategory
		return &t
	case protoobj.NotifyCategory_NotifyCategory_chats:
		t := typescore.ChatsMessageNotifyCategory
		return &t
	case protoobj.NotifyCategory_NotifyCategory_info:
		t := typescore.InfoNotifyCategory
		return &t
	case protoobj.NotifyCategory_NotifyCategory_temporary_password:
		t := typescore.TemporaryPasswordNotifyCategory
		return &t
	case protoobj.NotifyCategory_NotifyCategory_device_new:
		t := typescore.DeviceNewNotifyCategory
		return &t

	default:
		return nil
	}
}
