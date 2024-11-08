package useraccountmodule

import (
	protoobj "cnpc_backend/core/proto"
	"userservice/types"
)

func NewUserAccountServiceProto(ipc *types.InternalProviderControl) protoobj.UserServiceServer {
	return &UserAccountServiceProto{
		ipc: ipc,
	}
}

type UserAccountServiceProto struct {
	ipc *types.InternalProviderControl
	protoobj.UnimplementedUserServiceServer
}
