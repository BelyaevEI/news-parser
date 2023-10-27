package translater

import (
	"github.com/BelyaevEI/news-parser/internal/errors"
	"github.com/bregydoc/gtranslate"
)

var language = "ru"

// This function translate foreign text to Russian
func Translate(foreign string) (string, error) {

	translated, err := gtranslate.TranslateWithParams(
		foreign,
		gtranslate.TranslationParams{
			From: "auto",
			To:   language,
		},
	)
	if err != nil {
		return "", errors.Internal
	}
	return translated, nil
}
