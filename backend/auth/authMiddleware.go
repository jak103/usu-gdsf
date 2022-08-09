package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/jak103/usu-gdsf/config"
)

const (
	ACCESS_TOKEN_COOKIE_KEY = "accessToken"
	REFRESH_TOKEN_COOKIE_KEY = "refreshToken"
)

func RequireAuthorization(handler echo.HandlerFunc, requireAdmin bool) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessCookie, err := ctx.Cookie(ACCESS_TOKEN_COOKIE_KEY)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: No access token provided")
		}

		tokenClaims, err := DecodeAndVerifyToken(accessCookie.Value, ACCESS_TOKEN)

		if err != nil && err.Error() == TOKEN_EXPIRED {
			refreshCookie, err := ctx.Cookie(REFRESH_TOKEN_COOKIE_KEY)

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Token expired")
			}

			refreshTokenClaims, err := DecodeAndVerifyToken(refreshCookie.Value, REFRESH_TOKEN)

			if err != nil {
				if refreshTokenClaims.UserId != tokenClaims.UserId {
					return echo.NewHTTPError(http.StatusUnauthorized,
						fmt.Sprintf("Unauthorized: Could not refresh tokens: %s",
							"Refresh token and access tokens are for different users"))
				}

				return echo.NewHTTPError(http.StatusUnauthorized,
					fmt.Sprintf("Unauthorized: Could not refresh tokens: %s", err))
			}

			cookieLifetime := time.Duration((config.RefreshTokenLifetimeDays + 1) * 24) * time.Hour

			params := TokenParams{
				Type: ACCESS_TOKEN,
				UserId: tokenClaims.UserId,
				UserType: tokenClaims.UserType,
				UserEmail: tokenClaims.UserEmail,
			}
			
			newAccessCookie := http.Cookie{
				Name: ACCESS_TOKEN_COOKIE_KEY,
				Value: GenerateToken(params),
				Secure: true,
				HttpOnly: true,
				Expires: time.Now().Add(cookieLifetime),
			}

			params.Type = REFRESH_TOKEN
			
			newRefreshCookie := http.Cookie{
				Name: REFRESH_TOKEN_COOKIE_KEY,
				Value: GenerateToken(params),
				Secure: true,
				HttpOnly: true,
				Expires: time.Now().Add(cookieLifetime),
			}

			ctx.SetCookie(&newAccessCookie)
			ctx.SetCookie(&newRefreshCookie)
		} else if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err))
		}

		if requireAdmin && tokenClaims.UserType != ADMIN_USER {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized: Must be admin")
		}

		ctx.Set("tokenClaims", tokenClaims)

		return handler(ctx)
	}
}
