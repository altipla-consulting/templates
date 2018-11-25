package funcs

import (
	"time"

	"github.com/altipla-consulting/dateformatter"
	"github.com/altipla-consulting/datetime"
	pbdatetime "github.com/altipla-consulting/datetime/altipla/datetime"
	pbtimestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func DateFormat(t time.Time, lang, layout string) string {
	return dateformatter.Format(t, lang, layout)
}

func ProtoDateFormat(lang, layout string, t *pbdatetime.Date) string {
	return dateformatter.Format(datetime.ParseDate(t), lang, layout)
}

func TimestampFormat(lang, layout string, timestamp *pbtimestamp.Timestamp) string {
	return dateformatter.Format(datetime.ParseTimestamp(timestamp), lang, layout)
}
