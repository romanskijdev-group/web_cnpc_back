package utilscore

import (
	"cnpc_backend/core/typescore"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/big"
	"regexp"
	"strings"
	"time"
)

// GenerateTemporaryUserPassword - генерация временного пароля
func GeneratePassword(length *int, notSpecial bool) (*string, *typescore.WEvent) {
	if length == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("length is nil"),
			Text: "system_error",
		}
	}
	passwordBytes := make([]byte, *length)
	_, err := rand.Read(passwordBytes)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "system_error",
		}
	}
	passString := base64.StdEncoding.EncodeToString(passwordBytes)
	passString = strings.TrimRight(passString, "=")

	if notSpecial {
		reg := regexp.MustCompile("[^a-zA-Z0-9]+")
		passString = reg.ReplaceAllString(passString, "0")
	}
	return &passString, nil
}

// GenerateRandomCode - генерация случайного кода 6ти значного
func GenerateRandomCode() (*string, *typescore.WEvent) {
	maxValue := big.NewInt(999999)
	minValue := big.NewInt(100000)
	delta := new(big.Int).Sub(maxValue, minValue)
	delta = delta.Add(delta, big.NewInt(1))

	n, err := rand.Int(rand.Reader, delta)
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "system_error",
		}
	}

	codeGenerate := n.Add(n, minValue)
	codeGenerateStr := fmt.Sprintf("%d", codeGenerate)
	return &codeGenerateStr, nil
}

// сохранение временного секрета в Redis
func GenerateTemporarySecretUser(redisClient *redis.Client, userIdent string, lifeTime int, secretStr *string, typeSignature string) (*typescore.GenerateUserSecretI, *typescore.WEvent) {
	if secretStr == nil {
		return nil, &typescore.WEvent{
			Err:  fmt.Errorf("secretStr is nil"),
			Text: "system_error",
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userIdent = fmt.Sprintf("%s:%s", typeSignature, userIdent)

	_, err := redisClient.Get(ctx, userIdent).Result()
	if err == redis.Nil {
		// Код не найден, продолжаем и генерируем новый
		err := redisClient.Set(ctx, userIdent, *secretStr, time.Duration(lifeTime)*time.Minute).Err()
		if err != nil {
			log.Println("💔 error CheckAndDeleteFromRedis 2: ", err)
			return nil, &typescore.WEvent{
				Err:  errors.New("failed_to_set"),
				Text: "system_error",
			}
		}
	} else {
		// Код уже существует
		log.Println("💔 error CheckAndDeleteFromRedis 1: code already exists: ", err)
		return nil, &typescore.WEvent{
			Err:  errors.New("already_exists"),
			Text: "system_error",
		}
	}

	expiresAt := time.Now().UTC().Add(time.Duration(lifeTime) * time.Minute).Unix()

	return &typescore.GenerateUserSecretI{
		UserIdent: userIdent,
		Secret:    *secretStr,
		ExpiresIn: &expiresAt,
	}, nil
}
