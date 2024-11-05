package grpcclients

import (
	grpccore "cnpc_backend/core/grpc_core/grpc"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func InitClientUserAccountServiceProto(opts []grpc.DialOption, configObj *typescore.Config) protoobj.UserServiceClient {
	var conn *grpc.ClientConn
	var err error
	list := fmt.Sprintf("%s:%d", configObj.Server.UserService.Internal, configObj.Server.UserService.Port)
	// Попытка подключения с повторением
	for {
		conn, err = grpccore.CreateClientConnects(opts, list, false) // todo разобратся с чеком сервиса
		if err != nil {
			log.Printf("🔴 Failed to connect to UserAccountService server: %v. Retrying...", err)
			time.Sleep(1 * time.Second) // Задержка перед следующей попыткой
			continue
		}
		// Вывод сообщения о подключении только после успешного соединения
		log.Println("🟢 UserAccountService server connected... ", list)
		break // Выход из цикла, если подключение успешно
	}
	return protoobj.NewUserServiceClient(conn)
}
