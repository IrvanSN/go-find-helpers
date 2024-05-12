package routes

import (
	controllers "github.com/irvansn/go-find-helpers/controllers/thumbnail"
	"github.com/irvansn/go-find-helpers/middlewares"
	"github.com/labstack/echo/v4"
)

type ThumbnailRouteController struct {
	ThumbnailController *controllers.ThumbnailController
}

func (r *ThumbnailRouteController) InitRoute(e *echo.Echo) {
	c := e.Group("/v1/thumbnails")
	c.Use(middlewares.JWTMiddleware)
	c.GET("/upload-url", r.ThumbnailController.GetPreSignedURL)
}
