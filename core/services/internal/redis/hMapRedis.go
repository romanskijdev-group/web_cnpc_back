package redismodule

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"reflect"
	"strings"
	"time"
)

// обрабатывает ошибки, форматируя сообщение об ошибке.
func errorProcess(namesFunc string, mapHash string, key *string, err error) error {
	errStrInfo := fmt.Sprintf("⛔️ error: %s: failed to %s redis_value:", namesFunc, mapHash)
	if key != nil {
		errStrInfo = fmt.Sprintf("⛔️ %s: key: %s", errStrInfo, *key)
	}
	if err != nil {
		return fmt.Errorf("⛔️ %s: %v", errStrInfo, err)
	}
	return err
}

// получает значение из Redis hash map по ключу и десериализует его в obj.
func (m *ModuleRedis) GetHMapValueRedis(mapNames RedisHMapNames, key string, obj any) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Преобразование mapNames к типу string
	mapNamesStr := string(mapNames)

	result, err := m.GetClient().HGet(ctx, mapNamesStr, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, errorProcess("HGet GetHMapValueRedis", mapNamesStr, &key, err)
	}

	err = json.Unmarshal([]byte(result), obj)
	if err != nil {
		return nil, errorProcess("Unmarshal GetHMapValueRedis", mapNamesStr, &key, err)
	}

	return obj, nil
}

// сериализует obj в JSON и устанавливает его в Redis hash map по ключу.
func (m *ModuleRedis) SetHMapValueRedis(ctx context.Context, mapNames RedisHMapNames, key string, obj any, timeLiveMinute *int64) error {
	ctxRedisPersonal, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Преобразование mapNames к типу string
	mapNamesStr := string(mapNames)

	jsonObj, err := json.Marshal(obj)
	if err != nil {
		return errorProcess("jsonObj SetHMapValueRedis", mapNamesStr, &key, err)
	}

	if timeLiveMinute != nil {
		// Установить время жизни для ключа в минутах
		go func() {
			// Преобразовать timeLiveMinute в Duration
			duration := time.Duration(*timeLiveMinute) * time.Minute
			// Ожидать завершения таймера
			time.Sleep(duration)
			// Действие по завершении таймера
			// Например, удалить ключ из Redis
			err := m.DeleteHMapValueRedis(mapNames, key)
			if err != nil {
				// Обработка ошибки
				fmt.Println("Error deleting key:", key, " ", mapNames, err)
			}
		}()
		return nil
	}

	err = m.GetClient().HSet(ctxRedisPersonal, mapNamesStr, key, jsonObj).Err()
	if err != nil {
		return errorProcess("HSet SetHMapValueRedis", mapNamesStr, &key, err)
	}

	return nil
}

// удаляет значение из Redis hash map по ключу.
func (m *ModuleRedis) DeleteHMapValueRedis(mapNames RedisHMapNames, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Преобразование mapNames к типу string
	mapNamesStr := string(mapNames)

	err := m.GetClient().HDel(ctx, mapNamesStr, key).Err()
	if err != nil {
		return errorProcess("HDel DeleteHMapValueRedis", mapNamesStr, &key, err)
	}
	return nil
}

// получает все значения из Redis hash map, десериализует их и возвращает как список.
func (m *ModuleRedis) ListHMapValueRedis(mapNames RedisHMapNames, emptyObj any) ([]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Преобразование mapNames к типу string
	mapNamesStr := string(mapNames)

	result, err := m.GetClient().HGetAll(ctx, mapNamesStr).Result()
	if err != nil {
		return nil, errorProcess("HGetAll ListHMapValueRedis", mapNamesStr, nil, err)
	}
	objList := make([]any, 0, len(result))
	for _, value := range result {
		newObj := reflect.New(reflect.TypeOf(emptyObj).Elem()).Interface()
		err = json.Unmarshal([]byte(value), &newObj)
		if err != nil {
			errs := errorProcess("HGetAll ListHMapValueRedis", mapNamesStr, nil, err)
			log.Println(errs)
			continue
		}
		objList = append(objList, newObj)
	}

	return objList, nil
}

// удаляет все значения из Redis hash map.
func (m *ModuleRedis) ClearListHMapValueRedis(mapNames RedisHMapNames) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Преобразование mapNames к типу string
	mapNamesStr := string(mapNames)

	err := m.GetClient().Del(ctx, mapNamesStr).Err()
	if err != nil {
		return errorProcess("Del ClearListHMapValueRedis", mapNamesStr, nil, err)
	}
	return nil
}

// приведение ключей HMap Redis Для разных таблиц
func FormationLowKey(key0 string) string {
	return strings.ToLower(key0)
}

func Formation2ValLowlyKey(val0 string, val1 string) string {
	val0 = strings.ToLower(val0)
	val1 = strings.ToLower(val1)
	return fmt.Sprintf("%s::%s", val0, val1)
}
