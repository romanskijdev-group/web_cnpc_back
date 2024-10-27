package redismodule

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
	"time"
)

var (
	onceRedisModule     sync.Once
	redisModuleInstance *ModuleRedis
)

type ModuleRedis struct {
	RedisConfig *RedisConfig
	RedisClient *redis.Client
}

func GetRedisModuleInstance(redisConfigObj *RedisConfig) *ModuleRedis {
	onceRedisModule.Do(func() {
		redisClient := CreateRedisClient(redisConfigObj, 0)
		redisModuleInstance = &ModuleRedis{
			RedisConfig: redisConfigObj,
			RedisClient: redisClient,
		}
	})
	return redisModuleInstance
}

func CreateRedisClient(redisConfigObj *RedisConfig, redisNDb int) *redis.Client {
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisUser := ""
	redisPassword := ""
	redisConnectURI := ""

	if redisHost == "" || redisPort == "" {
		redisHost = redisConfigObj.Host
		redisPort = redisConfigObj.Port
	}
	if redisConfigObj.User != "" {
		redisUser = redisConfigObj.User
	}
	if redisConfigObj.Password != "" {
		redisPassword = redisConfigObj.Password
	} else {
		if redisConfigObj.ConnectURI != "" {
			redisConnectURI = redisConfigObj.ConnectURI
		} else {
			redisConnectURI = fmt.Sprintf("%s:%s", redisHost, redisPort)
		}
	}

	var opts *redis.Options
	if redisConnectURI != "" {
		opts = &redis.Options{
			Addr:     redisConnectURI,
			Username: redisUser,
			Password: redisPassword,
			DB:       redisNDb,
		}
	} else { // Обычное подключение
		opts = &redis.Options{
			Addr:         fmt.Sprintf("%s:%s", redisHost, redisPort),
			Username:     redisUser,
			Password:     redisPassword,
			DB:           redisNDb,
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     20,
			PoolTimeout:  30 * time.Second,
		}
	}
	return redis.NewClient(opts)
}

func (m *ModuleRedis) GetClient() *redis.Client {
	return m.RedisClient
}
