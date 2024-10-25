package typescore

import (
	"strings"
	"time"
)

type CustomDate time.Time

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	// Удаляем кавычки из строки
	str := strings.Trim(string(b), "\"")
	// Парсим строку в формат времени
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}
