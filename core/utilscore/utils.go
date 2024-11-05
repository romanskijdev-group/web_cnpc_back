package utilscore

import (
	"cnpc_backend/core/typescore"
	"errors"
	"github.com/mozillazg/go-unidecode"
	"net/mail"
	"runtime"
	"strings"
)

func ValidateEmailFormat(email *string) *typescore.WEvent {
	if email == nil || *email == "" {
		return &typescore.WEvent{
			Err:  errors.New("empty email"),
			Text: "empty_email",
		}
	}
	_, err := mail.ParseAddress(*email)
	if err != nil {
		return &typescore.WEvent{
			Err:  err,
			Text: "invalid_email_format",
		}
	}
	return nil
}

func Contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func OptimalGoroutines() int64 {
	return int64(runtime.NumCPU())
}

// FormatString Функция для форматирования строки
func FormatString(text *string) string {
	str := *text

	// Приведение к нижнему регистру
	str = strings.ToLower(str)

	// Транскрипция кириллических символов в латиницу
	str = unidecode.Unidecode(str)

	// Замена пробелов на нижние подчеркивания
	str = strings.ReplaceAll(str, " ", "_")
	str = strings.ReplaceAll(str, "'", "")

	return str
}
