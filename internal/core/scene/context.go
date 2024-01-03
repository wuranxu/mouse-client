// variable context
// use text/template

package scene

import (
	"bytes"
	"github.com/google/uuid"
	"text/template"
)

var tmpl = template.New("mouse").Funcs(map[string]any{
	"uuid": GetUuid,
})

func GetUuid(prefix string) string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return prefix + newUUID.String()
}

// Render render template's variables, like jinja2
func Render(origin string, variables map[string]any) (string, error) {
	tmpl, err := tmpl.Parse(origin)
	if err != nil {
		return origin, err
	}
	var b []byte
	buffer := bytes.NewBuffer(b)
	if err := tmpl.Execute(buffer, variables); err != nil {
		return origin, err
	}
	return buffer.String(), nil
}
