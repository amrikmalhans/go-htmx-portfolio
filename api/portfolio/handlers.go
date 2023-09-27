package portfolio

import (
	"html/template"
	"net/http"
)

func getPortfolio(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/index.html"))
	temp.Execute(w, nil)
}

func getHobbies(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/hobbies.html"))
	temp.Execute(w, nil)
}

func getWork(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/work.html"))
	temp.Execute(w, nil)
}

func getSchedule(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/schedule.html"))
	temp.Execute(w, nil)
}


func getJournals(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("web/templates/journals.html"))
	temp.Execute(w, nil)
}