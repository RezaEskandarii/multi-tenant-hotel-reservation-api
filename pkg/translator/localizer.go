package translator

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"strings"
)

type TranslateService interface {
	Localize(lang string, key string) string
}

type Translator struct {
	bundle i18n.Bundle
}

var (
	defaultLang = "en"
)

func New() *Translator {

	bundle := *i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if _, err := bundle.LoadMessageFile("resources/translations/en.json"); err != nil {
		panic(err.Error())
	}

	if _, err := bundle.LoadMessageFile("resources/translations/fa.json"); err != nil {
		panic(err.Error())
	}

	return &Translator{bundle: bundle}
}

// Localize Translates the given message into the given language.
func (t *Translator) Localize(lang string, key string) string {

	if lang == "" || strings.TrimSpace(lang) == "" {
		lang = defaultLang
	}

	loc := i18n.NewLocalizer(&t.bundle, lang)

	msg, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})

	if err != nil || strings.Trim(msg, " ") == "" {
		return ""
	}

	return msg
}
