package application

import (
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

	"github.com/vartzy/order-api-microservice/controller"
	"github.com/vartzy/order-api-microservice/repository/order"
)

func (a *App) loadRoutes() {
    router := chi.NewRouter()

    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    router.Route("/orders", a.loadOrderRoutes)

    a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {
    orderController := &controller.Order{
        Repo: &order.RedisRepo{
            Client: a.rdb,
        },
    }

    router.Get("/", orderController.List)
    router.Post("/", orderController.Create)
    router.Get("/{id}", orderController.GetById)
    router.Put("/{id}", orderController.UpdateById)
    router.Delete("/{id}", orderController.DeleteById)
}
