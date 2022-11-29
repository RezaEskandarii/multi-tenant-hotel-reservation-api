// Package translator localize messages /**/
package translator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"reservation-api/internal/global_variables"
	"strings"
)

var (
	defaultLang = "en"
	bundle      = *i18n.NewBundle(language.English)
)

func init() {

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if _, err := bundle.LoadMessageFile("resources/translations/en.json"); err != nil {
		panic(err.Error())
	}

	if _, err := bundle.LoadMessageFile("resources/translations/fa.json"); err != nil {
		panic(err.Error())
	}

}

// Localize Translates the given message into the given language.
func Localize(ctx context.Context, key string) string {

	langValue := ctx.Value(global_variables.CurrentLang)
	lang := ""

	if langValue != nil {
		lang = fmt.Sprintf("%s", langValue)
	}

	if lang == "" || strings.TrimSpace(lang) == "" {
		lang = defaultLang
	}

	loc := i18n.NewLocalizer(&bundle, lang)

	msg, err := loc.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})

	if err != nil || strings.Trim(msg, " ") == "" {
		return ""
	}

	return msg
}
