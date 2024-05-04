package middlewares

import (
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/irvansn/go-find-helpers/constant"
	"github.com/irvansn/go-find-helpers/controllers/base"
	"github.com/irvansn/go-find-helpers/utils"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
)

type Claims struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	Role      string
	jwt2.RegisteredClaims
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			err := constant.ErrNotAuthorized
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt2.ParseWithClaims(tokenString, &Claims{}, func(token *jwt2.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
				err := constant.ErrInvalidRequest
				return nil, c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
			}

			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			err := constant.ErrNotAuthorized
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		if !token.Valid {
			err := constant.ErrNotAuthorized
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			err := constant.ErrNotAuthorized
			return c.JSON(utils.ConvertResponseCode(err), base.NewErrorResponse(err.Error()))
		}

		c.Set("claims", claims)

		return next(c)
	}
}
