package controllers

import (
	"go-api-mini-shop/domain"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UseCase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{
		UseCase: usecase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var user domain.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Register(user)

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

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var user domain.User

	if err := ctx.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Login(user)

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
		Data:    map[string]string{"token": result},
	})
}
