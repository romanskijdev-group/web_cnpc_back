package marshallerglobal

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

// Serialization
func FilteringParamsListSerialization(obj *typescore.FilteringParamsList) *protoobj.FilteringParamsList {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.FilteringParamsList{
		Offset:         s.Uint64ToWrapperUInt64Value(obj.Offset),
		Limit:          s.Uint64ToWrapperUInt64Value(obj.Limit),
		LikeFields:     obj.LikeFields,
		OrSearchFields: obj.OrSearchFields,
	}
}
