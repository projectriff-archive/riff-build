package initializer

import (
	"bytes"
	"text/template"
)

func ApplyTemplate(tmpl string, name string, params interface{}) (string, error) {
	var buffer bytes.Buffer
	t, err := template.New(name).Parse(tmpl)
	if err == nil {
		err = t.Execute(&buffer, params)
	}
	return buffer.String(), err
}
