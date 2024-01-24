package controllers

import (
	"go-api-mini-shop/domain"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	UseCase domain.CategoryUsecase
}

func NewCategoryController(usecase domain.CategoryUsecase) *CategoryController {
	return &CategoryController{
		UseCase: usecase,
	}
}

func (c *CategoryController) GetAll(ctx *fiber.Ctx) error {
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

func (c *CategoryController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := c.UseCase.Detail(id)

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

func (c *CategoryController) Insert(ctx *fiber.Ctx) error {
	var category domain.Category

	if err := ctx.BodyParser(&category); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Insert(category.Name)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(domain.Response{
		Status:  fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	var category domain.Category

	if err := ctx.BodyParser(&category); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := c.UseCase.Update(category.ID, category.Name)

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

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
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
