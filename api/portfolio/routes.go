package portfolio

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PortfolioRoutes(r *chi.Mux) {
	// server static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	r.Route("/", func(r chi.Router) {
		r.Get("/", getPortfolio)
	})

	r.Route("/hobbies", func(r chi.Router) {
		r.Get("/", getHobbies)
	})

	r.Route("/work", func(r chi.Router) {
		r.Get("/", getWork)
	})

	r.Route("/schedule", func(r chi.Router) {
		r.Get("/", getSchedule)
	})

}

func JournalsRoutes(r *chi.Mux) {
	r.Route("/journals", func(r chi.Router) {
		r.Get("/", getJournals)
	})

	r.Route("/journals/{journal}", func(r chi.Router) {
		r.Get("/", getJournal)
	})
}
