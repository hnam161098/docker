package main

import (
	"docker/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := route.Route{
		App: fiber.New(),
	}
	server.RouteInit()

	server.App.Listen(":3000")
}
