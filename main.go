package main

import (
	"github.com/irvansn/go-find-helpers/config"
	addressController "github.com/irvansn/go-find-helpers/controllers/address"
	categoryController "github.com/irvansn/go-find-helpers/controllers/category"
	jobController "github.com/irvansn/go-find-helpers/controllers/job"
	ratingController "github.com/irvansn/go-find-helpers/controllers/rating"
	thumbnailController "github.com/irvansn/go-find-helpers/controllers/thumbnail"
	transactionController "github.com/irvansn/go-find-helpers/controllers/transaction"
	userController "github.com/irvansn/go-find-helpers/controllers/user"
	"github.com/irvansn/go-find-helpers/drivers/postgresql"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/address"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/category"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/job"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/thumbnail"

	"github.com/irvansn/go-find-helpers/drivers/postgresql/rating"

	"github.com/irvansn/go-find-helpers/drivers/postgresql/transaction"
	"github.com/irvansn/go-find-helpers/drivers/postgresql/user"
	"github.com/irvansn/go-find-helpers/routes"
	"github.com/irvansn/go-find-helpers/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.InitConfigPostgresql()
	db := postgresql.ConnectDB(config.InitConfigPostgresql())

	e := echo.New()
	e.Use(middleware.Logger())

	userRepo := user.NewUserRepo(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	newUserController := userController.NewUserController(userUseCase)

	categoryRepo := category.NewCategoryRepo(db)
	categoryUseCase := usecases.NewCategoryUseCase(categoryRepo)
	newCategoryController := categoryController.NewCategoryController(categoryUseCase)

	jobRepo := job.NewJobRepo(db)
	jobUseCase := usecases.NewJobUseCase(jobRepo)
	newJobController := jobController.NewJobController(jobUseCase)

	addressRepo := address.NewAddressRepo(db)
	addressUseCase := usecases.NewAddressUseCase(addressRepo)
	newAddressController := addressController.NewAddressController(addressUseCase)

	transactionRepo := transaction.NewTransactionRepo(db)
	transactionUseCase := usecases.NewTransactionUseCase(transactionRepo)
	newTransactionController := transactionController.NewTransactionController(transactionUseCase)

	thumbnailRepo := thumbnail.NewThumbnailRepo(db)
	thumbnailUseCase := usecases.NewThumbnailUseCase(thumbnailRepo)
	newThumbnailController := thumbnailController.NewThumbnailController(thumbnailUseCase)

	ratingRepo := rating.NewRatingRepo(db)
	ratingUseCase := usecases.NewRatingUseCase(ratingRepo)
	newRatingController := ratingController.NewRatingController(ratingUseCase)

	userRouteController := routes.UserRouteController{
		UserController: newUserController,
	}
	categoryRouteController := routes.CategoryRouteController{
		CategoryController: newCategoryController,
	}
	jobRouteController := routes.JobRouteController{
		JobController: newJobController,
	}
	addressRouteController := routes.AddressRouteController{
		AddressController: newAddressController,
	}
	transactionRouteController := routes.TransactionRouteController{
		TransactionController: newTransactionController,
	}
	thumbnailRouteController := routes.ThumbnailRouteController{
		ThumbnailController: newThumbnailController,
	}
	ratingRouteController := routes.RatingRouteController{
		RatingController: newRatingController,
	}

	userRouteController.InitRoute(e)
	categoryRouteController.InitRoute(e)
	jobRouteController.InitRoute(e)
	addressRouteController.InitRoute(e)
	transactionRouteController.InitRoute(e)
	thumbnailRouteController.InitRoute(e)
	ratingRouteController.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
