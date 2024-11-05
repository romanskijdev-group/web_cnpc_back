package loader

import (
	"fmt"
	"html/template"
	"log"
	"notification_service/types"
	"os"
)

func LoadMailTemplates() *types.TemplatesMailSystem {
	templatesMailObj := &types.TemplatesMailSystem{}
	basePath := "loader/mail-template"
	mailTemplatesNameMap := map[string]string{
		"TfaEmailTemplate":      "tfa-email.html",
		"MailLoginPassTemplate": "mail-login-pass.html",
		"NewDeviceInfoTemplate": "new-device-info.html",
		"CustomNotifyTemplate":  "custom-notify.html",
	}

	for key, value := range mailTemplatesNameMap {
		pathTemplate := fmt.Sprintf("%s/%s", basePath, value)
		if _, err := os.Stat(pathTemplate); err != nil { // Проверка существования файла
			log.Println("🔴 Failed to find mail template:", key, err)
			continue // Пропускаем шаблон, если не найден
		}
		fileBytes, err := os.ReadFile(pathTemplate) // Используем os.ReadFile
		if err != nil {
			log.Println("🔴 Failed to load mail template:", key, err)
			continue // Пропускаем шаблон, если не удалось загрузить
		}
		t, err := template.New(key).Parse(string(fileBytes))
		if err != nil {
			log.Println("🔴 Failed to load mail template:", key, err)
			continue // Пропускаем шаблон, если не удалось разобрать
		}
		switch key {
		case "TfaEmailTemplate":
			templatesMailObj.TfaEmailTemplate = t
		case "MailLoginPassTemplate":
			templatesMailObj.MailLoginPassTemplate = t
		case "NewDeviceInfoTemplate":
			templatesMailObj.NewDeviceInfoTemplate = t
		case "CustomNotifyTemplate":
			templatesMailObj.CustomNotifyTemplate = t
		}
	}
	return templatesMailObj
}
