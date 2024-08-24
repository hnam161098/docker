package route

import (
	"docker/controller"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	App *fiber.App
}

func (r *Route) RouteInit() {
	r.registerAPI()
}

func (r *Route) registerAPI() {
	v := r.App.Group("/v1.0")

	// customer
	customerGroup := v.Group("/customer")
	handler := controller.NewCustomerHandler()
	customerGroup.Post("/create", handler.CreateCustomerAPI)
	customerGroup.Get("/list", handler.FindCustomerAPI)
}
