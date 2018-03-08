package funcs

import (
	"html/template"

	"github.com/juju/errors"
)

func SafeHtml(s ...string) (template.HTML, error) {
	if len(s) == 0 {
		return template.HTML(""), nil
	}

	if len(s) > 1 {
		return template.HTML(""), errors.New("can only sanitize one content at a time")
	}

	return template.HTML(s[0]), nil
}

func SafeJs(s ...string) (template.JS, error) {
	if len(s) == 0 {
		return template.JS(""), nil
	}

	if len(s) > 1 {
		return template.JS(""), errors.New("can only sanitize one content at a time")
	}

	return template.JS(s[0]), nil
}

func SafeURL(s ...string) (template.URL, error) {
	if len(s) == 0 {
		return template.URL(""), nil
	}

	if len(s) > 1 {
		return template.URL(""), errors.New("can only sanitize one content at a time")
	}

	return template.URL(s[0]), nil
}
