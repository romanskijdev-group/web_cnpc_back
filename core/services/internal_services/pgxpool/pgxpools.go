package pgxpool

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type ConfigConnectPgxPool struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func ConnectDB(configObj *ConfigConnectPgxPool) *pgxpool.Pool {
	logrus.Info("🟨 ConnectDB")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configObj.Host,
		configObj.Port,
		configObj.User,
		configObj.Password,
		configObj.Name)

	configDB, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("📛 error: failed to parse config: ", err)
		return nil
	}

	// настройки пула
	configDB.MaxConns = 50                        // Максимальное количество открытых соединений в пуле
	configDB.MinConns = 5                         // Минимальное количество открытых соединений, которое пул будет поддерживать
	configDB.MaxConnLifetime = time.Second * 45   // Максимальное время жизни соединения. После этого времени соединение будет закрыто
	configDB.MaxConnIdleTime = time.Second * 45   // Максимальное время простоя соединения. Если соединение не используется в течение этого времени, оно будет закрыто
	configDB.HealthCheckPeriod = time.Second * 15 // Периодичность проверки состояния соединений в пуле

	pool, err := pgxpool.NewWithConfig(context.Background(), configDB)
	if err != nil {
		log.Fatal("📛 error: failed to connect to database: ", err)
		return nil
	}

	return pool
}
