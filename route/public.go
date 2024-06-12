package route

import (
    "github.com/go-chi/chi/v5"

    "github.com/vartzy/order-api-microservice/controller"
)

func LoadOrderRoutes(router chi.Router) {
    orderController := &controller.Order{}

    router.Get("/", orderController.List)
    router.Post("/", orderController.Create)
    router.Get("/{id}", orderController.GetById)
    router.Put("/{id}", orderController.UpdateById)
    router.Delete("/{id}", orderController.DeleteById)
}
