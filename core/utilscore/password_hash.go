package utilscore

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Функция для хэширования пароля
func HashPassword(password *string) (*string, error) {
	if password == nil {
		return nil, errors.New("missed password")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	hashedPass := string(bytes)
	return &hashedPass, err
}

// Функция для проверки пароля
func CheckPasswordHash(password, hash *string) bool {
	if password == nil || hash == nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
