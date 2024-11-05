package interfacenotification

import (
	"gopkg.in/gomail.v2"
)

type MailSystemActionsI interface {
	TFANotifyMail(objBodyParams map[string]string) (*gomail.Message, error)
	TemporaryPasswordNotifyMail(objBodyParams map[string]string) (*gomail.Message, error)
	DeviceNewInfoNotifyMail(objBodyParams map[string]string) (*gomail.Message, error)
	CustomNotifyMail(objBodyParams map[string]string) (*gomail.Message, error)
}
