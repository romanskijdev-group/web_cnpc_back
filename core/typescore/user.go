package typescore

import (
	"time"
)

// user Role - роли пользователей
type UserRoleTypes string

const (
	UserRole       UserRoleTypes = "user"        // пользователь
	AdminRole      UserRoleTypes = "admin"       // администратор
	SuperAdminRole UserRoleTypes = "super_admin" // супер администратор
	SupportRole    UserRoleTypes = "support"     // поддержка
	MerchantRole   UserRoleTypes = "merchant"    // мерчант
)

type VerifiedKYCLevel uint8

const (
	NotVerifiedKYCLevel             VerifiedKYCLevel = 0 // Не верифицирован
	NameAndDocumentVerifiedKYCLevel VerifiedKYCLevel = 1 // Верифицированы имя и документ
	FullVerifiedKYCLevel            VerifiedKYCLevel = 2 // Полностью верифицирован
)

// UsersProviderControl - структура для управления пользователями + данные пользователя
type UsersProviderControl struct {
	SystemID                *string        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:system_id" json:"system_id,omitempty" db:"system_id"`       // Системный идентификатор записи
	SerialID                *uint64        `gorm:"type:bigint;index;autoIncrement;unique;column:serial_id" json:"serial_id" db:"serial_id" mapstructure:"serial_id"` // Уникальный порядковый идентификатор записи
	Role                    *UserRoleTypes `gorm:"type:varchar(20);index;default:'user';column:role" json:"role" db:"role"`                                          // Роль пользователя
	Email                   *string        `gorm:"type:varchar(255);index;unique;column:email" json:"email,omitempty" db:"email" mapstructure:"email"`               // Адрес электронной почты пользователя
	TelegramID              *int64         `gorm:"unique;index;column:telegram_id" json:"telegram_id" db:"telegram_id"`                                              // Идентификатор пользователя в Telegram
	Nickname                *string        `gorm:"type:varchar(50);index;unique;column:nickname" json:"nickname,omitempty" db:"nickname" mapstructure:"nickname"`    // Псевдоним или никнейм пользователя
	FirstName               *string        `gorm:"type:varchar(50);column:first_name" json:"first_name,omitempty" db:"first_name"`                                   // Имя пользователя
	LastName                *string        `gorm:"type:varchar(50);column:last_name" json:"last_name,omitempty" db:"last_name"`                                      // Фамилия пользователя
	Bio                     *string        `gorm:"type:text;column:bio" json:"bio,omitempty" db:"bio"`                                                               // Фамилия пользователя
	Gender                  *string        `gorm:"type:varchar(10);column:gender" json:"gender,omitempty" db:"gender"`                                               // Пол пользователя
	BirthDate               *time.Time     `gorm:"type:timestamp with time zone;column:birth_date" json:"birth_date" db:"birth_date"`                                // Дата рождения пользователя
	PhoneNumber             *uint64        `gorm:"unique;column:phone_number" json:"phone_number" db:"phone_number"`                                                 // Номер телефона пользователя
	AvatarURL               *string        `gorm:"type:varchar(255);column:avatar_url" json:"avatar_url,omitempty" db:"avatar_url"`                                  // URL аватара пользователя
	Language                *string        `gorm:"type:varchar(5);default:'en';column:language" json:"language,omitempty" db:"language"`                             // Язык настроек пользователя
	PushNotificationEnabled *bool          `gorm:"default:true;column:push_notification_enabled" json:"push_notification_enabled" db:"push_notification_enabled"`    // Включены ли разрешения на push-уведомления
	IsBlocked               *bool          `gorm:"default:false;column:is_blocked" json:"is_blocked" db:"is_blocked"`                                                // Залочен ли пользователь(заблокирован или нет)

	EmailVerified *bool      `gorm:"default:false;column:email_verified" json:"email_verified" db:"email_verified"`           // Подтвержден ли адрес электронной почты
	ReferralID    *string    `gorm:"type:uuid;column:referral_id" json:"referral_id,omitempty" db:"referral_id"`              // Идентификатор реферала пригласившего пользователя
	ReferralCode  *string    `gorm:"type:varchar(10);column:referral_code" json:"referral_code,omitempty" db:"referral_code"` // Реферальный код
	LastIP        *string    `gorm:"type:text;column:last_ip" json:"last_ip" db:"last_ip"`                                    // Последний IP-адрес входа пользователя
	CreatedAt     *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at" db:"created_at"`           // Дата и время создания записи
	LastLogin     *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:last_login" json:"last_login" db:"last_login"`           // Дата последнего входа пользователя

	IsOnline   *bool      `gorm:"default:false;column:is_online" json:"is_online" db:"is_online"`                   //Онлайн ли пользователь
	LastOnline *time.Time `gorm:"default:CURRENT_TIMESTAMP;column:last_online" json:"last_online" db:"last_online"` // Дата и время последнего онлайна
}

// укороченная информация о пользователе
type ShortUserInfo struct {
	SerialID   *uint64 `json:"serial_id,omitempty" db:"serial_id"`     // Уникальный порядковый идентификатор записи
	Email      *string `json:"email,omitempty" db:"email"`             // Адрес электронной почты пользователя
	TelegramID *int64  `json:"telegram_id,omitempty" db:"telegram_id"` // Идентификатор пользователя в Telegram
	Nickname   *string `json:"nickname,omitempty" db:"nickname"`       // Псевдоним или никнейм пользователя
	FirstName  *string `json:"first_name,omitempty" db:"first_name"`   // Имя пользователя
	LastName   *string `json:"last_name,omitempty" db:"last_name"`     // Фамилия пользователя
	ParentName *string `json:"parent_name,omitempty" db:"parent_name"` // Отчество пользователя
	AvatarURL  *string `json:"avatar_url,omitempty" db:"avatar_url"`   // URL аватара пользователя

	VerificationP2P  *bool             `json:"verification_p2p" db:"verification_p2p"`     // Подтвержден ли пользователь в P2P(выдается администратором)
	VerifiedKYCLevel *VerifiedKYCLevel `json:"verified_kyc_level" db:"verified_kyc_level"` // Уровень верификации KYC пользователя
	P2PTradeStatus   *bool             `json:"p2p_trade_status" db:"p2p_trade_status"`     // Статус торговли в P2P
	CreatedAt        *string           `json:"created_at" db:"created_at"`                 // Дата и время создания аккаунта

	IsOnline   *bool      `json:"is_online" db:"is_online"`     //Онлайн ли пользователь
	LastOnline *time.Time `json:"last_online" db:"last_online"` // Дата и время последнего онлайна
}
