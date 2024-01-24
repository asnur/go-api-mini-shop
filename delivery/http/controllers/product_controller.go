package controllers

import (
	"go-api-mini-shop/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	UseCase domain.ProductUsecase
}

func NewProductController(usecase domain.ProductUsecase) *ProductController {
	return &ProductController{
		UseCase: usecase,
	}
}

func (c *ProductController) GetAll(ctx *fiber.Ctx) error {
	result, err := c.UseCase.GetAll()

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

func (c *ProductController) GetByID(ctx *fiber.Ctx) error {
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

func (c *ProductController) Insert(ctx *fiber.Ctx) error {
	var product domain.Product

	if err := ctx.BodyParser(&product); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Create(product)

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

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	var product domain.Product

	if err := ctx.BodyParser(&product); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Update(product)

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

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
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
		Data:    nil,
	})
}
