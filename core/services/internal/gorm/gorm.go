package gormmodule

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigConnectGorm struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func connectGormDB(configObj *ConfigConnectGorm) (*gorm.DB, error) {
	logrus.Info("🟨 connectGormDB")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configObj.Host,
		configObj.Port,
		configObj.User,
		configObj.Password,
		configObj.Name)

	// Настройка логгера GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // медленный SQL порог
			LogLevel:                  logger.Silent, // уровень логирования
			IgnoreRecordNotFoundError: true,          // игнорировать ошибки записи не найдены
			Colorful:                  false,         // отключить цвет
		},
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Подключение к базе данных GORM
func GormDatabaseConnect(configObj *ConfigConnectGorm) (*gorm.DB, *sql.DB) {
	logrus.Info("🟨 GormDatabaseConnect")
	if configObj == nil {
		panic("configObj is nil")
	}
	db, err := connectGormDB(configObj)
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns устанавливает максимальное количество соединений в пуле простаивающих соединений.
	sqlDB.SetMaxIdleConns(1)

	// SetMaxOpenConns устанавливает максимальное количество открытых соединений с базой данных.
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime устанавливает максимальное время, в течение которого может быть повторно использовано соединение.
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		panic("failed to get DB from GORM")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return db, sqlDB
}
