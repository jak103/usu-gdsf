package auth

import (
        "net/http"
        "net/http/httptest"
        "testing"

        "github.com/labstack/echo/v4"
        "github.com/stretchr/testify/assert"
)

const(
        HANDLER_STUB_SUCCESS = "Success!"
)

func handlerStub(ctx echo.Context) error {
        return ctx.String(http.StatusOK, HANDLER_STUB_SUCCESS)
}


func TestMiddlewareWithValidToken(t *testing.T) {
        e := echo.New()
        rec := httptest.NewRecorder()
        req := httptest.NewRequest(http.MethodGet, "/_testStub", nil)
        c := e.NewContext(req, rec)

        params := TokenParams{
                Type: ACCESS_TOKEN,
                UserId: 39,
                UserType: REGULAR_USER,
                UserEmail: "testing@example.com",
        }

        req.Header.Set("Cookie", ACCESS_TOKEN_COOKIE_KEY + "=" + GenerateToken(params))
        stubRequest := RequireAuthorization(handlerStub, false)

        if assert.NoError(t, stubRequest(c)) {
                respBody := rec.Body.String()
                assert.Equal(t, respBody, HANDLER_STUB_SUCCESS)
        }
}


