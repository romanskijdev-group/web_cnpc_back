package marshallerusersubscription

import (
	marshallerutils "zod_backend_dev/core/grpc_core/marshaller_utils"
	"zod_backend_dev/core/models"
	protoobj "zod_backend_dev/core/proto"
)

func UserSubscriptionReqSerialization(obj *models.UsersSubscriptions) *protoobj.UsersSubscriptionsMsg {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UsersSubscriptionsMsg{
		SerialId:               s.Uint64ToWrapperUInt64Value(obj.SerialID),
		UserId:                 s.StringToWrapperStringValue(obj.UserID),
		SubscriptionName:       s.StringToWrapperStringValue(obj.SubscriptionName),
		CompatibilitiesBalance: s.Int32ToWrapperInt32Value(obj.CompatibilitiesCount),
		PersonalHoroscopeCount: s.Int32ToWrapperInt32Value(obj.PersonalHoroscopeCount),

		IsRenewal: s.BoolToWrapperBoolValue(obj.IsRenewal),
		StartDate: s.TimePtrToTimestampPB(obj.StartDate),
		ExpiredIn: s.TimePtrToTimestampPB(obj.ExpiredIn),
	}
}
