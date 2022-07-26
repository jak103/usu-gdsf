package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RequireAuthorization(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := ctx.Request().Header.Get("Authorization")
		splitHeader := strings.Split(authorizationHeader, " ")
		
		if len(splitHeader) != 2 || strings.ToLower(splitHeader[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Invalid Authorization header")
		}
		
		tokenClaims, err := DecodeAndVerifyToken(splitHeader[1], ACCESS_TOKEN)
		
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
		}

		ctx.Set("tokenClaims", tokenClaims)

		return handler(ctx)
	}
}

