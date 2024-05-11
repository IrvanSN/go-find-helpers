package thumbnail

import (
	"fmt"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/labstack/echo/v4"
)

type ThumbnailController struct {
	thumbnailUseCase entities.ThumbnailUseCaseInterface
}

func (tc *ThumbnailController) GetPreSignedURL(c echo.Context) error {
	// only CUSTOMER user can access
	fileName := c.QueryParam("file_name")

	fmt.Println(fileName)

	return nil
}

func NewThumbnailController(thumbnailUseCase entities.ThumbnailUseCaseInterface) *ThumbnailController {
	return &ThumbnailController{
		thumbnailUseCase: thumbnailUseCase,
	}
}
