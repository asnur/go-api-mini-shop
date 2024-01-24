package routes

import (
	"go-api-mini-shop/delivery/http/controllers"
	"go-api-mini-shop/domain"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Controllers struct {
	CartController     controllers.CartController
	UserController     controllers.UserController
	CateGoryController controllers.CategoryController
	ProductController  controllers.ProductController
	OrderController    controllers.OrderController
}

func RegisterRoutes(c Controllers, ctx *fiber.App) {
	// Add Middleware CORS
	ctx.Use(cors.New())

	// Grouping V1
	v1 := ctx.Group("/api/v1")

	// User
	userRoutes := v1.Group("/user")
	userRoutes.Post("/register", c.UserController.Register)
	userRoutes.Post("/login", c.UserController.Login)

	// Restricted Routes
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(domain.Response{
				Status:  fiber.StatusUnauthorized,
				Message: "Unauthorized",
			})
		},
	}))

	// Category
	categoryRoutes := v1.Group("/category")
	categoryRoutes.Get("/", c.CateGoryController.GetAll)
	categoryRoutes.Get("/:id", c.CateGoryController.GetByID)
	categoryRoutes.Post("/", c.CateGoryController.Insert)
	categoryRoutes.Put("/:id", c.CateGoryController.Update)
	categoryRoutes.Delete("/:id", c.CateGoryController.Delete)

	// Product
	productRoutes := v1.Group("/product")
	productRoutes.Get("/", c.ProductController.GetAll)
	productRoutes.Get("/:id", c.ProductController.GetByID)
	productRoutes.Post("/", c.ProductController.Insert)
	productRoutes.Put("/:id", c.ProductController.Update)
	productRoutes.Delete("/:id", c.ProductController.Delete)

	// Cart
	cartRoutes := v1.Group("/cart")
	cartRoutes.Get("/", c.CartController.GetAll)
	cartRoutes.Post("/", c.CartController.Insert)
	cartRoutes.Put("/", c.CartController.Update)
	cartRoutes.Delete("/:id", c.CartController.Delete)

	// Order
	orderRoutes := v1.Group("/order")
	orderRoutes.Get("/", c.OrderController.GetAll)
	orderRoutes.Get("/:id", c.OrderController.Detail)
	orderRoutes.Post("/", c.OrderController.Insert)
	orderRoutes.Delete("/:id", c.OrderController.Delete)
}
