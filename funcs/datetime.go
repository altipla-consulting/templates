package funcs

import (
	"time"

	"github.com/altipla-consulting/dateformatter"
	"github.com/altipla-consulting/datetime"
	pbdatetime "github.com/altipla-consulting/datetime/protos/datetime"
)

func DateFormat(t time.Time, lang, layout string) string {
	return dateformatter.Format(t, lang, layout)
}

func ProtoDateFormat(lang, layout string, t *pbdatetime.Date) string {
	return dateformatter.Format(datetime.ParseDate(t), lang, layout)
}
