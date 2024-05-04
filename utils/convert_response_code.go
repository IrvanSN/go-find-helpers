package utils

import (
	"errors"
	"github.com/irvansn/go-find-helpers/constant"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch {
	case errors.Is(err, constant.ErrInsertDatabase):
		return http.StatusInternalServerError
	case errors.Is(err, constant.ErrEmptyInput):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrDuplicatedData):
		return http.StatusBadRequest
	case errors.Is(err, constant.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, constant.ErrInvalidEmailOrPassword):
		return http.StatusUnauthorized
	case errors.Is(err, constant.ErrNotAuthorized):
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
