package module

import (
	"cnpc_backend/core/typescore"
	"context"
	"errors"
	"log"
	"strings"
	"time"
	"userservice/locale"
)

func (s *UserAccountServiceProto) sendLoginAlertNotification(userInfo *typescore.UsersProviderControl, authType *typescore.TypeAuth) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if authType == nil {
		return errors.New("unexpected_auth_type")
	}

	replaceMapText := map[string]string{
		"<<service>>": string(*authType),
	}

	translateTitle, err := locale.LocaleConvert(userInfo.Language, "title_new_login", s.ipc.Modules.BundleI18n)
	if err != nil {
		log.Println("🔴 error LocaleConvert: ", err)
		translateTitle = ""
	}
	translateBody, err := locale.LocaleConvert(userInfo.Language, "body_new_login", s.ipc.Modules.BundleI18n)
	if err != nil {
		log.Println("🔴 error LocaleConvert: ", err)
		translateBody = ""
	}

	bodyText := translateBody
	for key, value := range replaceMapText {
		bodyText = strings.ReplaceAll(bodyText, key, value)
	}

	notifyType := typescore.InfoNotifyCategory

	userAlert := &typescore.UserSystemAlerts{
		UserID:     userInfo.SystemID,
		NotifyType: &notifyType,
		Title:      &translateTitle,
		Message:    &bodyText,
	}

	_, errW := s.ipc.Database.UserAlerts.CreateUserAlertDB(ctx, userAlert)
	if errW != nil {
		return errW.Err
	}

	return nil
}
