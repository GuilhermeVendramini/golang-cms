package config

import (
	"html/template"
)

// TPL template.Template
var TPL *template.Template

func init() {
	TPL = template.Must(template.New("").ParseGlob("templates/*.html"))
}
