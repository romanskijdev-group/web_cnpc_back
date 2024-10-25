package grpccore

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// создает соединение с gRPC сервером
func CreateClientConnects(opts []grpc.DialOption, baseUrl string, healthCheck bool) (*grpc.ClientConn, error) {
	// logrus.Info("🟨 CreateClientConnects")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	log.Println("🟡 Connecting to gRPC server... ", baseUrl)
	maxMessageSize := 250 * 1024 * 1024                                                   // Установка максимального размера сообщения в 250 МБ
	maxSizeOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMessageSize)) // Создание опции с максимальным размером сообщения
	opts = append(opts, maxSizeOption)                                                    // Добавление опции с максимальным размером сообщения в список опций

	// Установка соединения с gRPC сервером с использованием указанных опций
	conn, err := grpc.NewClient(baseUrl, opts...)
	if err != nil {
		return nil, err // Возвращение ошибки, если соединение не установлено
	}

	if healthCheck {
		healthClient := healthpb.NewHealthClient(conn)
		ctxHealth, cancelHealth := context.WithTimeout(context.Background(), time.Second)
		defer cancelHealth()

		resp, err := healthClient.Check(ctxHealth, &healthpb.HealthCheckRequest{
			Service: "", // Укажите имя сервиса, если нужно проверить конкретный сервис, или оставьте пустым для проверки сервера в целом
		})
		if err != nil {
			log.Println(" 🔴 error healthClient: ", err)
			return nil, err
		}

		log.Printf("Состояние сервиса: %s", resp.Status)

		// Ожидание установки соединения
		for {
			state := conn.GetState()
			if state == connectivity.Ready {
				break // Выход из цикла, если соединение готово
			}
			if !conn.WaitForStateChange(ctx, state) {
				return nil, ctx.Err() // Возвращение ошибки, если ожидание изменения состояния было прервано
			}
			time.Sleep(1 * time.Second) // Задержка перед следующей проверкой состояния соединения
		}
	}

	return conn, nil
}
