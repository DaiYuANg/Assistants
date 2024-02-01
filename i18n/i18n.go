package i18n

import (
	"embed"
	"github.com/BurntSushi/toml"
	"github.com/gohugoio/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locale.*.toml
var LocaleFS embed.FS

func newI18n() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile("es.toml")
	if err != nil {
		return
	}
}
