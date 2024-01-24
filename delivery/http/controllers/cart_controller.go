package controllers

import (
	"go-api-mini-shop/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type CartController struct {
	UseCase domain.CartUsecase
}

func NewCartController(usecase domain.CartUsecase) *CartController {
	return &CartController{
		UseCase: usecase,
	}
}

func (c *CartController) GetAll(ctx *fiber.Ctx) error {
	// Get User ID from JWT
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	params := map[string]interface{}{
		"user_id": int(claims["id"].(float64)),
	}

	result, err := c.UseCase.GetAll(params)

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

func (c *CartController) Insert(ctx *fiber.Ctx) error {
	var cart domain.Cart

	if err := ctx.BodyParser(&cart); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// Get User ID from JWT
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	cart.UserID = int(claims["id"].(float64))

	result, err := c.UseCase.Insert(&cart)

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

func (c *CartController) Update(ctx *fiber.Ctx) error {
	var cart domain.Cart

	if err := ctx.BodyParser(&cart); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Update(cart.ID, cart.Quantity)

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

func (c *CartController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
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
		Data:    nil,
	})
}
