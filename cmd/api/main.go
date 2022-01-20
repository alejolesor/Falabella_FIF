package main

import (
	"net/http"

	"falabella/cmd/api/app"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := httplog.NewLogger("http-log", httplog.Options{
		Concise: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))

	router(r)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func router(r *chi.Mux) {
	beers := app.DependencyInjection()

	r.Get("/beers", beers.Get)
	r.Get("/beers/{beerID}", beers.GetBeerxId)
	r.Post("/beers", beers.Post)
	r.Get("/beers/{beerID}/boxprice", beers.BoxPrice)
}
