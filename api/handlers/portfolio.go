package handlers

import (
	"html/template"
	"net/http"
)

func GetPortfolio(w http.ResponseWriter, _ *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/index.html"))
	temp.Execute(w, nil)
}
