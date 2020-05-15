package models

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type IConfig interface {
	InitializeConfig(input interface{})
}

var DefaultLanguage *Language

func init() {
	DefaultLanguage = &Language{
		AcceptLanguage: "en",
		Localizer:      i18n.NewLocalizer(i18n.NewBundle(language.English), "en"),
	}
}

func SetDefaultLanguage(language *Language) {
	DefaultLanguage = language
}
