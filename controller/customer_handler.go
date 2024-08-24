package controller

import (
	"context"
	"docker/repositories"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerRepo repositories.CustomerModel
}

func NewCustomerHandler() *CustomerHandler {
	return &CustomerHandler{
		customerRepo: repositories.NewCustomerModel(),
	}
}

func (c *CustomerHandler) CreateCustomerAPI(ctx *fiber.Ctx) error {
	var name = fmt.Sprintf("nam_%v", time.Now().Unix())
	var age = 100
	var address = "Hanoi"
	model := repositories.CustomerModel{
		Name:      name,
		Age:       int64(age),
		Address:   address,
		CreatedAt: time.Now(),
	}
	errC := c.customerRepo.InsertOneCustomer(context.Background(), model)
	if errC != nil {
		msg := fmt.Sprintf("errors: %s", errC)
		return ctx.JSON(msg)
	}
	return ctx.JSON("Create customer success")
}

func (c *CustomerHandler) FindCustomerAPI(ctx *fiber.Ctx) error {
	result, err := c.customerRepo.FindCustomer(context.Background())
	if err != nil {
		msg := fmt.Sprintf("errors: %s", err)
		return ctx.JSON(msg)
	}
	return ctx.JSON(result)
}
