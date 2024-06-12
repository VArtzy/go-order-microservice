package route

import (
	"github.com/go-chi/chi/v5"

	"github.com/vartzy/order-api-microservice/controller"
	"github.com/vartzy/order-api-microservice/repository/order"
    "github.com/vartzy/order-api-microservice/types"
)

func (a *types.App) LoadOrderRoutes(router chi.Router) {
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
