package typescore

import (
	"net"
	"time"
)

// Структура для доступа к API
type APIAccess struct {
	SerialID     *uint64    `gorm:"type:bigint;autoIncrement;unique;column:serial_id" json:"serial_id" db:"serial_id"`         // Уникальный порядковый идентификатор записи
	CreatedAt    *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at" db:"created_at"`             // Дата и время создания записи
	APIPublic    *string    `gorm:"type:varchar(255);not null;primaryKey;column:api_public" json:"api_public" db:"api_public"` // Ключ доступа к API
	APIPrivate   *string    `gorm:"type:varchar(255);not null;column:api_private" json:"api_private" db:"api_private"`         // Секретный ключ доступа к API
	IsActive     *bool      `gorm:"default:true;column:is_active" json:"is_active" db:"is_active"`                             // Статус доступа к API
	Desctription *string    `gorm:"type:text;column:description" json:"description" db:"description"`                          // Описание доступа к API
	TypeAccess   *string    `gorm:"type:varchar(255);column:type_access;default:'user'" json:"type_access" db:"type_access"`   // Тип доступа к API
}

// Черный список ip
type BlackListIP struct {
	IP        *net.IP    `gorm:"type:inet;primaryKey;column:ip" json:"ip" db:"ip"`                              // Ip адрес в черном списке
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at" db:"created_at"` // Дата и время создания записи
}
