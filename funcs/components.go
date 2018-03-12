package funcs

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
)

type componentInstance struct {
	name   string
	params []*componentParam
}

type componentParam struct {
	name  string
	value interface{}
}

func Component(name string) *componentInstance {
	return &componentInstance{name: name}
}

func EndComponent(c *componentInstance) (template.HTML, error) {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, `<div data-vue="%s"`, c.name)
	for _, param := range c.params {
		switch v := param.value.(type) {
		case proto.Message:
			m := jsonpb.Marshaler{EmitDefaults: true}
			b, err := m.MarshalToString(v)
			if err != nil {
				return template.HTML(""), errors.Trace(err)
			}
			param.name = fmt.Sprintf("$%v", param.name)
			param.value = html.EscapeString(b)

		case bool, int32, int, float64:
			param.name = fmt.Sprintf("$%v", param.name)

		case string:
			param.value = html.EscapeString(v)

		case int64:
			param.value = fmt.Sprintf("%v", v)

		case []int32:
			var s []string
			for _, n := range v {
				s = append(s, fmt.Sprintf("%d", n))
			}
			param.name = fmt.Sprintf("$%v", param.name)
			param.value = fmt.Sprintf("[%s]", strings.Join(s, ","))
		}

		fmt.Fprintf(&buf, ` data-%v="%v"`, param.name, param.value)
	}
	fmt.Fprintf(&buf, `></div>`)

	return template.HTML(buf.String()), nil
}

func ComponentParam(c *componentInstance, name string, value interface{}) string {
	c.params = append(c.params, &componentParam{name, value})
	return ""
}
