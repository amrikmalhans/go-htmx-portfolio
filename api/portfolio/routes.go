package portfolio

import "github.com/go-chi/chi/v5"

func PortfolioRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", getPortfolio)
	})
}
