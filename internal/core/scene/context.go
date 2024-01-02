// variable context
// use text/template

package scene

import (
	"bytes"
	"text/template"
)

var tmpl = template.New("mouse")

func Render(origin string, variables map[string]any) (string, error) {
	tmpl, err := tmpl.Parse(origin)
	if err != nil {
		return origin, err
	}
	var ans []byte
	buffer := bytes.NewBuffer(ans)
	if err := tmpl.Execute(buffer, variables); err != nil {
		return origin, err
	}
	return buffer.String(), nil
}
