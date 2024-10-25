package marshallerglobal

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

// Deserialization

func FilteringParamsListDeserialization(obj *protoobj.FilteringParamsList) *typescore.FilteringParamsList {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.FilteringParamsList{
		Offset:         d.OptionalUint64(obj.Offset),
		Limit:          d.OptionalUint64(obj.Limit),
		LikeFields:     obj.LikeFields,
		OrSearchFields: obj.OrSearchFields,
	}
}
