package utils

import (
	"bytes"
	"html/template"
)

func Parse(t *template.Template, name string, data interface{}) (content string, err error) {
	var doc bytes.Buffer
	if err = t.ExecuteTemplate(&doc, name, data); err != nil {
		return
	}
	content = doc.String()
	return
}
