package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/transaction"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type TransactionRouteController struct {
	TransactionController *controllers.TransactionController
}

func (r *TransactionRouteController) InitRoute(e *echo.Echo) {
	t := e.Group("/v1/transactions")
	t.Use(middlewares.JWTMiddleware)
	t.GET("", r.TransactionController.GetAllTransactions)
}
