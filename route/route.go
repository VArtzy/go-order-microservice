package route

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes() *chi.Mux {
    router := chi.NewRouter()

    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    router.Route("/orders", LoadOrderRoutes)

    return router
}