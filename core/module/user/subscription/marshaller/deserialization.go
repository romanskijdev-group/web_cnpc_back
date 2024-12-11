package marshallerusersubscription

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func UserSubscriptionReqDeserialization(obj *protoobj.UsersSubscriptionsMsg) *typescore.UsersSubscriptions {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.UsersSubscriptions{
		SerialID:         d.OptionalUint64(obj.SerialId),
		UserID:           d.OptionalString(obj.UserId),
		SubscriptionName: d.OptionalString(obj.SubscriptionName),
		StartDate:        d.OptionalTime(obj.StartDate),
		ExpiredIn:        d.OptionalTime(obj.ExpiredIn),
		IsRenewal:        d.OptionalBool(obj.IsRenewal),
	}
}
