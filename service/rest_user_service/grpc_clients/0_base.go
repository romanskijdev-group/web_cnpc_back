package grpcclients

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateDialOptionsProto() []grpc.DialOption {
	// Определение опций подключения (без специфических опций)
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return opts
}
