package redismodule

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

func (m *ModuleRedis) GetRedisMemoryUsage() (*RedisMemoryInfo, error) {
	// logrus.Info("🟨 GetRedisMemoryUsage")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := m.RedisClient.Info(ctx, "memory").Result()
	if err != nil {
		log.Println("🔴 error GetRedisMemoryUsage: ", err)
		return nil, errors.New("not_found")
	}

	fmt.Println("🔰🔰 Memory usage:", result)
	lines := strings.Split(result, "\n")
	info := &RedisMemoryInfo{}
	v := reflect.ValueOf(info).Elem()

	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		fieldName := strings.ReplaceAll(parts[0], "-", "")
		field := v.FieldByName(fieldName)
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			field.SetString(parts[1])
		}
	}

	return info, nil
}
