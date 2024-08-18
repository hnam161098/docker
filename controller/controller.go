package controller

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ControllerHandler struct {
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{}
}

func (c *ControllerHandler) SayHello(ctx *fiber.Ctx) error {
	t := time.Now().Format("15:04:03 02/01/2006")
	msg := fmt.Sprintf("Bây giờ là %v", t)
	return ctx.JSON(msg)
}
