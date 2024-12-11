package typescore

import "time"

// UsersProviderControl - структура для управления пользователями + данные пользователя
type UserSystemAlerts struct {
	SystemID  *string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:system_id" json:"system_id,omitempty" db:"system_id" mapstructure:"system_id"` // Системный идентификатор записи
	SerialID  *uint64    `gorm:"type:bigint;index;autoIncrement;unique;column:serial_id" json:"serial_id" db:"serial_id" mapstructure:"serial_id"`                    // Уникальный порядковый идентификатор записи
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at" db:"created_at"`                                                       // Дата и время создания записи

	UserID     *string         `gorm:"type:uuid;column:user_id" json:"user_id,omitempty" db:"user_id" mapstructure:"user_id"`                                        // Системный идентификатор пользователя
	Reading    *bool           `gorm:"type:boolean;column:reading;default:false" json:"reading,omitempty" db:"reading" mapstructure:"reading"`                       // Флаг прочтения уведомления пользователем
	NotifyType *NotifyCategory `gorm:"type:varchar(100);index;not null;column:notify_type" json:"notify_type,omitempty" db:"notify_type" mapstructure:"notify_type"` //  Тип уведомления
	Title      *string         `gorm:"type:varchar(100);column:title" json:"title,omitempty" db:"title"`                                                             // Заголовок уведомления
	Message    *string         `gorm:"type:varchar(500);column:message" json:"message,omitempty" db:"message"`                                                       // Сообщение уведомления
	Link       *string         `gorm:"type:text;column:link" json:"link,omitempty" db:"link"`                                                                        // Ссылка на дополнительную информацию
	DeepLinkID *uint64         `gorm:"type:bigint;column:deep_link_id" json:"deep_link_id,omitempty" db:"deep_link_id"`                                              // Идентификатор глубокой ссылки
}
