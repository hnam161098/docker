package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var env = os.Getenv("ENV_ARGS")

type ControllerHandler struct {
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{}
}

func (c *ControllerHandler) SayHello(ctx *fiber.Ctx) error {
	t := time.Now().Format("15:04:03 02/01/2006")
	msg := fmt.Sprintf("Đây là môi trường: %s và bây giờ là %v", env, t)
	return ctx.JSON(msg)
}
