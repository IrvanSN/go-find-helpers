package main

import (
	"github.com/irvansn/go-find-helpers/config"
	categoryController "github.com/irvansn/go-find-helpers/controllers/category"
	userController "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/irvansn/go-find-helpers/drivers/postgresql"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
	"github.com/irvansn/go-find-helpers/routes"
	"github.com/irvansn/go-find-helpers/usecases"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitConfigMySQL()
	db := postgresql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	userRepo := user.NewUserRepo(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	newUserController := userController.NewUserController(userUseCase)

	categoryRepo := category.NewCategoryRepo(db)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)
	newCategoryController := categoryController.NewCategoryController(categoryUseCase)

	userRouteController := routes.UserRouteController{
		UserController: newUserController,
	}
	categoryRouteController := routes.CategoryRouteController{
		CategoryController: newCategoryController,
	}

	userRouteController.InitRoute(e)
	categoryRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
