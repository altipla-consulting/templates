package funcs

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/altipla-consulting/content"
	"github.com/altipla-consulting/messageformat"
	"github.com/juju/errors"
)

var messages = map[string]map[string]string{}

func init() {
	f, err := os.Open("messages.json")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}

		panic(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&messages); err != nil {
		panic(err)
	}
}

func NativeName(lang string) string {
	names := map[string]string{
		"es": "Español",
		"ca": "Català",
		"en": "English",
		"fr": "Français",
	}
	return names[lang]
}

func MsgFormat(lang, format string, params ...interface{}) (string, error) {
	format = Translate(lang, format)

	message, err := messageformat.New(format)
	if err != nil {
		return "", errors.Trace(err)
	}

	res, err := message.Format(lang, params)
	if err != nil {
		return "", errors.Trace(err)
	}

	return res, nil
}

func Translate(lang, format string) string {
	if lang == "es" {
		return format
	}

	msg, ok := messages[format]
	if !ok {
		msg = map[string]string{}
	}

	// En producción se parte correctamente la descripción; pero en desarrollo
	// debemos quitarla partiendo el valor resultante de la cadena de lenguajes.
	return strings.Split(content.LangChain(msg, lang), "//")[0]
}
