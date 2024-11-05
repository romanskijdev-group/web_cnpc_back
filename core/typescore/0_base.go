package typescore

// MailType - типы писем
const (
	MailTypeTemporaryPassword = "mt_temporary_password" // временный пароль
	MailTypeNewDeviceAlert    = "mt_new_device_alert"   // уведомление о новом устройстве
)

type SortingParamsType string

const (
	DescSortingParams SortingParamsType = "desc" // По убыванию
	AscSortingParams  SortingParamsType = "asc"  // По возрастанию
)

// CurrencyTypes - типы валют
type CurrencyTypes string

const (
	Fiat   CurrencyTypes = "fiat"   // фиатные валюты
	Crypto CurrencyTypes = "crypto" // криптовалюты
)

// Добавочные параметры для хранения в REDIS
const (
	TempPassRedisType       = "rds_t_pass"              // временный пароль
	TfaCodeSignature        = "rds_sign_code"           // подпись кода для 2fa
	TempCodeTelegramBotType = "rds_t_code_telegram_bot" // временный код для телеграм бота
)

const (
	DBOperationCreate = "create" // создание
	DBOperationUpdate = "update" // обновление
	DBOperationDelete = "delete" // удаление
	DBOperationRead   = "read"   // чтение
)

type IsActivate struct {
	IsActivate *bool `json:"is_activate"`
}

type StructSignature struct {
	Signature *string `json:"signature"`
}

type FilteringMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
