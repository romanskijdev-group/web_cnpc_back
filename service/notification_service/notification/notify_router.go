package notification

import (
	"cnpc_backend/core/typescore"
	"errors"

	"gopkg.in/gomail.v2"
)

type MsgNotifyStruct struct {
	User      typescore.UsersProviderControl
	TitleText string
	BodyText  string

	MailMessage        *gomail.Message
	OnEmail            bool
	AlertNotifyAppType *typescore.NotifyCategory
}

func (m *ModuleNotification) NotifyRouting(notifyParams *typescore.NotifyParams) error {
	if notifyParams == nil {
		return errors.New("notifyParams is nil")
	}

	if notifyParams.Category == nil {
		return errors.New("notifyParams.Category is nil")
	}

	switch *notifyParams.Category {
	case typescore.TemporaryPasswordNotifyCategory: // Временный пароль
		return m.TemporaryPasswordNotifyCategoryAction(notifyParams)
	case typescore.DeviceNewNotifyCategory: // Новое устройство
		return m.DeviceNewNotifyCategoryAction(notifyParams)
	}
	return nil
}
