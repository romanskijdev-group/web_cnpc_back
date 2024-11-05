package typescore

import (
	"gopkg.in/gomail.v2"
)

// Определение типа функции для отправки уведомлений
type NotificationFunc func(objBodyParams map[string]string) (*gomail.Message, error)

type MailSMTPConfig struct {
	PortSMTP     string `json:"port_smtp"`
	HostSMTP     string `json:"host_smtp"`
	SenderSMTP   string `json:"sender_smtp"`
	PasswordSMTP string `json:"password_smtp"`
}

type MailResultI struct {
	Email     *string `json:"email"`
	ExpiresIn *int64  `json:"expires_in"`
}

type ParamsSendMail struct {
	MailType       string
	EmailRecipient string
	IPAddress      *string
}

// Тип/Категория уведомления

type NotifyCategory string

const (
	BearerChequeNotifyCategory NotifyCategory = "bearer_cheque" //  чеки
	ChatsMessageNotifyCategory NotifyCategory = "chats"         // Чаты
	InfoNotifyCategory         NotifyCategory = "info"          // Информационное уведомление(от админа)

	TemporaryPasswordNotifyCategory NotifyCategory = "temporary_password" // Временный пароль
	DeviceNewNotifyCategory         NotifyCategory = "device_new"         // Новое устройство
)

type NotifyAdditionsObject struct {
	//BearerChequeObj *BearerCheque // Чек(при необходимости)
	SubmittedBy *string // Кем отправлено(при необходимлсти)
}

type NotifyParams struct {
	Text          *string         // Текст уведомления
	Title         *string         // Заголовок уведомления
	SystemUserIDs []*string       // Список идентификаторов пользователей
	MailAddress   *string         // Адрес электронной почты
	Category      *NotifyCategory // Категория уведомления

	IsEmail   bool // Отправка уведомления на email
	Emergency bool // Экстренная отправка уведомления игнорирует запреты пользователя

	AdditionsObject *NotifyAdditionsObject // Дополнительные данные
}
