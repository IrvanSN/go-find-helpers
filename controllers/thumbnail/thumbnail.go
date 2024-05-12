package thumbnail

import (
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/controllers/thumbnail/response"
	"github.com/irvansn/go-find-helpers/entities"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ThumbnailController struct {
	thumbnailUseCase entities.ThumbnailUseCaseInterface
}

func (tc *ThumbnailController) GetPreSignedURL(c echo.Context) error {
	var getPreSignedURLRequest entities.Thumbnail
	fileName := c.QueryParam("file_name")
	getPreSignedURLRequest.ImageKey = "/thumbnail/" + fileName

	thumbnail, errUseCase := tc.thumbnailUseCase.GetPreSignedURL(&getPreSignedURLRequest)
	if errUseCase != nil {
		return c.JSON(utils.ConvertResponseCode(errUseCase), base.NewErrorResponse(errUseCase.Error()))
	}

	getPreSignedURLResponse := response.FromUseCase(&thumbnail)
	return c.JSON(http.StatusOK, base.NewSuccessResponse("PreSigned URL has created!", getPreSignedURLResponse))
}

func NewThumbnailController(thumbnailUseCase entities.ThumbnailUseCaseInterface) *ThumbnailController {
	return &ThumbnailController{
		thumbnailUseCase: thumbnailUseCase,
	}
}
