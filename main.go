package main

import (
	"github.com/irvansn/go-find-helpers/config"
	controllers "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/irvansn/go-find-helpers/drivers/mysql"
	"github.com/irvansn/go-find-helpers/drivers/mysql/user"
	"github.com/irvansn/go-find-helpers/routes"
	"github.com/irvansn/go-find-helpers/usecases"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	config.InitConfigMySQL()
	db := mysql.ConnectDB(config.InitConfigMySQL())

	e := echo.New()

	userRepo := user.NewUserRepo(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	routeController := routes.RouteController{
		UserController: userController,
	}

	routeController.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
