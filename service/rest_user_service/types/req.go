package types

import "time"

// запрос обновление профиля пользователя
type UpdateUserProfileHandlerReq struct {
	Nickname            *string    `json:"nickname,omitempty"`        // Псевдоним или никнейм пользователя
	FirstName           *string    `json:"first_name,omitempty"`      // Имя пользователя
	LastName            *string    `json:"last_name,omitempty"`       // Фамилия пользователя
	BirthDate           *time.Time `json:"birth_date"`                // Дата рождения пользователя
	PhoneNumber         *string    `json:"phone_number"`              // Номер телефона пользователя
	Language            *string    `json:"language,omitempty"`        // Язык настроек пользователя
	NotificationEnabled *bool      `json:"push_notification_enabled"` // Включены ли разрешения на push-уведомления
}

type AvatarURLSetter struct {
	URL string `json:"url"`
}
