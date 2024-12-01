package locale

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var LanguageMapTranslate = map[string]bool{
	"en": true,
	"ru": true,
}

func LocaleConvert(langCode *string, tag string, bundle *i18n.Bundle) (string, error) {
	langCodeStr := "en"
	if langCode != nil {
		langCodeStr = *langCode
		if !LanguageMapTranslate[langCodeStr] {
			langCodeStr = "en"
		}
	}
	localizer := i18n.NewLocalizer(bundle, langCodeStr)
	// Получаем перевод, используя localizer
	translation, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: tag})
	if err != nil {
		// В случае ошибки возвращаем неизмененную подстроку
		return "", err
	}
	return translation, nil
}
