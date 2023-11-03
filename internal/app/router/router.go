package router

import (
	"net/http"
	"refactoring/internal/app/controller"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(c controller.User) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", c.Search)
				r.Post("/", c.Create)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", c.Get)
					r.Patch("/", c.Update)
					r.Delete("/", c.Delete)
				})
			})
		})
	})

	return r
}