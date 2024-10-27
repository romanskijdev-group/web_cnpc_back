package marshallerusersubscription

import (
	marshallerutils "zod_backend_dev/core/grpc_core/marshaller_utils"
	"zod_backend_dev/core/models"
	protoobj "zod_backend_dev/core/proto"
)

func UserSubscriptionReqDeserialization(obj *protoobj.UsersSubscriptionsMsg) *models.UsersSubscriptions {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &models.UsersSubscriptions{
		SerialID:               d.OptionalUint64(obj.SerialId),
		UserID:                 d.OptionalString(obj.UserId),
		SubscriptionName:       d.OptionalString(obj.SubscriptionName),
		CompatibilitiesCount:   d.OptionalInt32(obj.CompatibilitiesBalance),
		PersonalHoroscopeCount: d.OptionalInt32(obj.PersonalHoroscopeCount),
		StartDate:              d.OptionalTime(obj.StartDate),
		ExpiredIn:              d.OptionalTime(obj.ExpiredIn),
		IsRenewal:              d.OptionalBool(obj.IsRenewal),
	}
}
