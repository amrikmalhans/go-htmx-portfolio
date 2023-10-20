package utils

import "text/template"

var Temp *template.Template

func InitTemplates() {
	// The order in which the templates are parsed is important
	// because of the way the base template is defined.
	// Better way to do this would be to use a glob pattern.
	Temp = template.Must(template.ParseFiles(
		"web/templates/base.html",
		"web/templates/components/navbar.html",
		"web/templates/journals.html",
		"web/templates/journal_slug.html",
	))
}
