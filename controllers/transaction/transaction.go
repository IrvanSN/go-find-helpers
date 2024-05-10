package transaction

import (
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/transaction/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionController struct {
	transactionUseCase entities.TransactionUseCaseInterface
}

func (tc *TransactionController) GetAllTransactions(c echo.Context) error {
	var getAllTransactionsRequest []entities.Transaction
	if err := c.Bind(&getAllTransactionsRequest); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
	}

	transactionType := c.QueryParam("type")

	userData, ok := c.Get("claims").(*middlewares.Claims)
	if !ok {
		return echo.ErrInternalServerError
	}

	transactions, errUseCase := tc.transactionUseCase.GetAllTransactions(&getAllTransactionsRequest, transactionType, userData)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	getAllTransactionsResponse := response.GetAllTransactions(&transactions)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Successfully delete user!", getAllTransactionsResponse))
}

func NewTransactionController(transactionUseCase entities.TransactionUseCaseInterface) *TransactionController {
	return &TransactionController{
		transactionUseCase: transactionUseCase,
	}
}
