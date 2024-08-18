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
	group := v.Group("/api")
	api := controller.NewControllerHandler()
	group.Get("/say_hello", api.SayHello)
}
