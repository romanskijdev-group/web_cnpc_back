package restauthcore

import (
	"cnpc_backend/core/typescore"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type TokenRestAuth struct {
	ConfigG *typescore.Config
}

// генерация JWT токена
func (m *TokenRestAuth) GetAuthToken(userObj *typescore.UsersProviderControl) (*typescore.TokenInfo, *typescore.WEvent) {

	jWTSecret := m.ConfigG.Secure.JWTSecret
	sessionHourLive := m.ConfigG.Secure.SessionTokenHoursLife
	expIn := time.Now().UTC().Add(time.Hour * time.Duration(sessionHourLive)).Unix()
	if userObj == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("user_obj_is_nil"),
			Text: "invalid_auth",
		}
	}
	if userObj.SystemID == nil || userObj.Role == nil {
		return nil, &typescore.WEvent{
			Err:  errors.New("user_obj_params_is_nil"),
			Text: "invalid_auth",
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  *userObj.SystemID,
		"role": *userObj.Role,
		"iat":  time.Now().UTC().Unix(),
		"exp":  expIn,
	})

	tokenStr, err := token.SignedString([]byte(jWTSecret))
	if err != nil {
		return nil, &typescore.WEvent{
			Err:  err,
			Text: "invalid_auth",
		}
	}
	return &typescore.TokenInfo{
		AccessToken: tokenStr,
		ExpiresIn:   expIn,
	}, nil
}

// получение данных с токена
func (m *TokenRestAuth) GetValuesFromToken(tokenStr string, jWTSecret string) (string, string, int64, *typescore.WEvent) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jWTSecret), nil
	})
	if err != nil {
		log.Println("🔴 error parse token: ", err)
		return "", "", 0, &typescore.WEvent{
			Err:  err,
			Text: "invalid_auth",
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			return "", "", 0, &typescore.WEvent{
				Err:  errors.New("invalid sub claim"),
				Text: "invalid_auth",
			}
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", "", 0, &typescore.WEvent{
				Err:  errors.New("invalid exp claim"),
				Text: "invalid_auth",
			}
		}

		role, ok := claims["role"].(string)
		if !ok {
			return "", "", 0, &typescore.WEvent{
				Err:  errors.New("invalid role claim"),
				Text: "invalid_auth",
			}
		}

		return sub, role, int64(exp), nil
	}
	return "", "", 0, &typescore.WEvent{
		Err:  errors.New("invalid token"),
		Text: "invalid_auth",
	}
}
