package api

import (
	"github.com/amrikmalhans/go-htmx-portfolio.git/api/handlers"
	"github.com/go-chi/chi/v5"
)

func PortfolioRoutes(r *chi.Mux) {

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetPortfolio)
	})
}

func JournalsRoutes(r *chi.Mux) {
	r.Route("/journals", func(r chi.Router) {
		r.Get("/", handlers.GetJournals)
	})

	r.Route("/journals/{journal}", func(r chi.Router) {
		r.Get("/", handlers.GetJournal)
	})
}
