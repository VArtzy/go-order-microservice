package route

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/vartzy/order-api-microservice/types"
)

func (a *types.App) LoadRoutes() {
    router := chi.NewRouter()

    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    router.Route("/orders", a.LoadOrderRoutes)

    a.router = router
}
