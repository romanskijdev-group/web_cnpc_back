package marshallerusersubscription

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func UserSubscriptionReqSerialization(obj *typescore.UsersSubscriptions) *protoobj.UsersSubscriptionsMsg {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UsersSubscriptionsMsg{
		SerialId:         s.Uint64ToWrapperUInt64Value(obj.SerialID),
		UserId:           s.StringToWrapperStringValue(obj.UserID),
		SubscriptionName: s.StringToWrapperStringValue(obj.SubscriptionName),

		IsRenewal: s.BoolToWrapperBoolValue(obj.IsRenewal),
		StartDate: s.TimePtrToTimestampPB(obj.StartDate),
		ExpiredIn: s.TimePtrToTimestampPB(obj.ExpiredIn),
	}
}
