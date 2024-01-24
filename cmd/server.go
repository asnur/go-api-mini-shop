package cmd

import (
	"go-api-mini-shop/config"
	"go-api-mini-shop/delivery/http/controllers"
	"go-api-mini-shop/delivery/http/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"

	category_repository "go-api-mini-shop/repository/category"
	category_usecase "go-api-mini-shop/usecase/category"

	product_repository "go-api-mini-shop/repository/product"
	product_usecase "go-api-mini-shop/usecase/product"

	user_repository "go-api-mini-shop/repository/user"
	user_usecase "go-api-mini-shop/usecase/user"
)

var (
	ip, port string

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run Server",
		Long:  `Run Server with IP and Port`,
		Run: func(cmd *cobra.Command, args []string) {

			// Connect to Database
			db, err := config.ConnectDB()

			if err != nil {
				panic(err)
			}

			// Set Address
			address := ip + ":" + port

			// Initalize Resources
			categoryRepository := category_repository.NewRepository(db)
			categoryUseCase := category_usecase.NewUseCase(categoryRepository)
			categoryController := controllers.NewCategoryController(categoryUseCase)

			productRepository := product_repository.NewRepository(db)
			productUseCase := product_usecase.NewUseCase(productRepository)
			productController := controllers.NewProductController(productUseCase)

			userRepository := user_repository.NewRepository(db)
			userUseCase := user_usecase.NewUseCase(userRepository)
			userController := controllers.NewUserController(userUseCase)

			// Initialize Controllers
			registeredController := &routes.Controllers{
				UserController:     *userController,
				CateGoryController: *categoryController,
				ProductController:  *productController,
			}

			app := fiber.New()

			// Middleware Logger
			app.Use(logger.New(logger.Config{
				Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
			}))

			// Register Routes
			routes.RegisterRoutes(*registeredController, app)

			// Run Server
			if err := app.Listen(address); err != nil {
				panic(err)
			}
		},
	}
)
