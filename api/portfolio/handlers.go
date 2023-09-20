package portfolio

import (
	"html/template"
	"net/http"
)

func getPortfolio(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/index.html"))
	temp.Execute(w, nil)
}
