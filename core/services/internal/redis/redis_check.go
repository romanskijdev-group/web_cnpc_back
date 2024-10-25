package redismodule

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Проверка и удаление временных данных данных из Redis
func (m *ModuleRedis) CheckAndDeleteFromRedis(userIdent, expectedValue, identType string) error {
	// logrus.Info("🟨 CheckAndDeleteFromRedis")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userIdentCustom := fmt.Sprintf("%s:%s", identType, userIdent)

	result, err := m.RedisClient.Get(ctx, userIdentCustom).Result()
	if err == redis.Nil {
		log.Println("💔 error CheckAndDeleteFromRedis 1: ", err)
		return errors.New("not_found")
	} else if err != nil {
		log.Println("💔 error CheckAndDeleteFromRedis 0: ", err)
		return errors.New("not_found")
	}

	if result != expectedValue {
		return errors.New("invalid_code")
	}

	_, err = m.RedisClient.Del(ctx, userIdentCustom).Result()
	if err != nil {
		log.Println("💔 error CheckAndDeleteFromRedis: ", err)
		return errors.New("not_found")
	}
	return nil
}
