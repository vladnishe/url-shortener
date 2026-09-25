package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "works fine")
	})

	return r
}
