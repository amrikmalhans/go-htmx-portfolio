package utils

import "text/template"

var Temp *template.Template

func InitTemplates() {
	Temp = template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/components/navbar.html",
		"web/templates/journal_slug.html",
		"web/templates/journals.html",
	))

}
