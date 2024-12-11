package typescore

import (
	"github.com/shopspring/decimal"
	"time"
)

type UsersSubscriptions struct {
	SerialID         *uint64 `gorm:"type:bigint;primaryKey;autoIncrement;column:serial_id" json:"serial_id" db:"serial_id" mapstructure:"serial_id"`                                                // Уникальный порядковый идентификатор записи
	UserID           *string `gorm:"type:uuid;not null;column:user_id" json:"user_id,omitempty" db:"user_id" mapstructure:"user_id"`                                                                // Системный идентификатор пользователя
	SubscriptionName *string `gorm:"type:varchar(100);not null;column:subscription_name;default:'user'" json:"subscription_name,omitempty" db:"subscription_name" mapstructure:"subscription_name"` // Название тарифа пользователя

	ProjectsCount   *int32 `gorm:"not null;column:projects_count;default:5" json:"projects_count" db:"projects_count"`       // Остатки количества совместимостей
	MaxContributors *int32 `gorm:"not null;column:max_contributors;default:3" json:"max_contributors" db:"max_contributors"` // Остатки количества инд. гороскопов

	IsRenewal *bool      `gorm:"type:bool;not null;default:true;column:is_renewal" json:"is_renewal" db:"is_renewal"`
	StartDate *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:start_date" json:"start_date" db:"start_date"` // Дата и время начала
	ExpiredIn *time.Time `gorm:"column:expired_in" json:"expired_in" db:"expired_in"`                           // Дата и время истекания
}

type Subscription struct {
	SystemID *string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:system_id" json:"system_id,omitempty" db:"system_id" mapstructure:"system_id"` // Системный идентификатор записи
	SerialID *uint64 `gorm:"type:bigint;index;autoIncrement;unique;column:serial_id" json:"serial_id" db:"serial_id"`                                             // Уникальный порядковый идентификатор записи

	Name            *string          `gorm:"type:varchar(100);not null;column:name" json:"name,omitempty" db:"name" mapstructure:"name"` // Название
	Description     *string          `gorm:"type:text;column:description" json:"description,omitempty" db:"description"`                 // Описание
	ProjectsCount   *int32           `gorm:"not null;column:projects_count;default:5" json:"projects_count" db:"projects_count"`         // Остатки количества совместимостей
	MaxContributors *int32           `gorm:"not null;column:max_contributors;default:3" json:"max_contributors" db:"max_contributors"`   // Остатки количества инд. гороскопов
	Duration        *int32           `gorm:"not null;column:duration" json:"duration" db:"duration" mapstructure:"duration"`             // Продолжительность (дней)
	Price           *decimal.Decimal `gorm:"not null;column:price" json:"price" db:"price"`                                              // Стоимость (usd)
	IsActive        *bool            `gorm:"not null;default:false;column:is_active" json:"is_active" db:"is_active"`                    // Активна ли подписка
	IsDefault       *bool            `gorm:"not null;default:false;column:is_default" json:"is_default" db:"is_default"`                 // Является ли тарифом по умолчанию
}
