package i18n

import (
	"embed"
)

//go:embed locale.*.toml
var LocaleFS embed.FS

func newI18n() {
	//bundle := i18n.NewBundle(language.English)
	//bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//_, err := bundle.LoadMessageFile("es.toml")
	//if err != nil {
	//	return
	//}
}
