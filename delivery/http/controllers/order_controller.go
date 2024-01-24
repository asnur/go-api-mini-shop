package controllers

import (
	"go-api-mini-shop/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type OrderController struct {
	UseCase domain.OrderUsecase
}

func NewOrderController(usecase domain.OrderUsecase) *OrderController {
	return &OrderController{
		UseCase: usecase,
	}
}

func (c *OrderController) GetAll(ctx *fiber.Ctx) error {
	// Get User ID from JWT
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := int(claims["id"].(float64))

	result, err := c.UseCase.GetAll(user_id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if len(result) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(domain.Response{
			Status:  fiber.StatusNotFound,
			Message: "Data not found",
			Data:    nil,
		})
	}

	return ctx.JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (c *OrderController) Insert(ctx *fiber.Ctx) error {
	var order struct {
		CartID []int `json:"cart_id"`
	}

	if err := ctx.BodyParser(&order); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Get User ID from JWT
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := int(claims["id"].(float64))

	result, err := c.UseCase.Insert(user_id, order.CartID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (c *OrderController) Detail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := c.UseCase.GetByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

func (c *OrderController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = c.UseCase.Delete(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(domain.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
	})
}
