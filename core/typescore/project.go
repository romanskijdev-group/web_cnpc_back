package typescore

import "time"

// Тип/Категория проекта

type ProjectCategory string

const (
	SystemProjectCategory ProjectCategory = "system" // системный
	UserProjectCategory   ProjectCategory = "user"   // пользовательский
)

// UserProject - структура для управления проектами
type UserProject struct {
	SystemID  *string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:system_id" json:"system_id,omitempty" db:"system_id" mapstructure:"system_id"` // Системный идентификатор записи
	SerialID  *uint64    `gorm:"type:bigint;index;autoIncrement;unique;column:serial_id" json:"serial_id" db:"serial_id" mapstructure:"serial_id"`                    // Уникальный порядковый идентификатор записи
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at" db:"created_at"`                                                       // Дата и время создания записи

	OwnerID     *string          `gorm:"type:uuid;column:owner_id" json:"owner_id,omitempty" db:"owner_id" mapstructure:"owner_id"`                                        // Системный идентификатор пользователя
	Name        *string          `gorm:"type:varchar(100);column:name" json:"name,omitempty" db:"name" mapstructure:"name"`                                                // Имя проекта
	Description *string          `gorm:"type:varchar(500);column:description" json:"description,omitempty" db:"description"`                                               // Описание проекта
	ProjectType *ProjectCategory `gorm:"type:varchar(100);index;not null;column:project_type" json:"project_type,omitempty" db:"project_type" mapstructure:"project_type"` // Тип проекта
	BannerURL   *string          `gorm:"type:varchar(500);column:banner_url" json:"banner_url,omitempty" db:"banner_url"`                                                  // Ссылка на баннер
	Attachment  *string          `gorm:"type:varchar(500);column:attachment" json:"attachment,omitempty" db:"attachment"`                                                  // Принадлежность
}
