package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amrikmalhans/go-htmx-portfolio.git/api"
	"github.com/amrikmalhans/go-htmx-portfolio.git/utils"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// init static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// init templates
	utils.InitTemplates()

	api.PortfolioRoutes(r)
	api.JournalsRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port, // for railway.app
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server gracefully stopped")
}
