package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RequireAuthorization(handler echo.HandlerFunc, requireAdmin bool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var token string

		if authCookie, err := ctx.Cookie("accessToken"); err == nil {
			token = authCookie.Value
		} else if authHeader := ctx.Request().Header.Get("Authorization"); authHeader != "" {
			splitHeader := strings.Split(authHeader, " ")
			
			if len(splitHeader) != 2 || strings.ToLower(splitHeader[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized,
					"Unauthorized: Invalid Authorization header")
			}

			token = splitHeader[1]
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: No access token provided")
		}
		
		tokenClaims, err := DecodeAndVerifyToken(token, ACCESS_TOKEN)
		
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
		}

		if requireAdmin && tokenClaims.UserType != ADMIN_USER {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Must be admin")
		}

		ctx.Set("tokenClaims", tokenClaims)

		return handler(ctx)
	}
}

