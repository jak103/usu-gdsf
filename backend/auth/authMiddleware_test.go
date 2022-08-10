// Test for expired token
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


func TestMiddlewareWithexpiredToken(t *testing.T) {
        e := echo.New()
        rec := httptest.NewRecorder()
        req := httptest.NewRequest(http.MethodGet, "/_testStub", nil)
        c := e.NewContext(req, rec)

        

        req.Header.Set("Cookie", ACCESS_TOKEN_COOKIE_KEY + "=eyJUeXBlIjowLCJFeHBpcmF0aW9uIjoxNjU5NDAyNTI5OTg1LCJVc2VySWQiOjQyLCJVc2VyVHlwZSI6MSwiVXNlckVtYWlsIjoidGVzdGluZ0BleGFtcGxlLmNvbSJ9fGE5N2QwMjViYzhkNzc2YjQwZjkyNmExNTBkMzg3OWEyY2Q5YTNmZTVkYTA1YTIyM2VlMDRmMDZlNGQxODFkZGI")
        stubRequest := RequireAuthorization(handlerStub, false)

        assert.Error(t, stubRequest(c))
//         if assert.NOError(t, stubRequest(c)) {
//                 respBody := rec.Body.String()
//                 assert.Equal(t, respBody, HANDLER_STUB_SUCCESS)
//         }
 }


// Provide access token cookie that is valid and not expired
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

//Test to not provide an access token

func TestMiddlewareNoToken(t *testing.T) {
        e := echo.New()
        rec := httptest.NewRecorder()
        req := httptest.NewRequest(http.MethodGet, "/_testStub", nil)
        c := e.NewContext(req, rec)

        stubRequest := RequireAuthorization(handlerStub, false)
        assert.Error(t, stubRequest(c))
}


//Provide access token cookie that is INVALID

func TestMiddlewareWitheinvalidToken(t *testing.T) {
        e := echo.New()
        rec := httptest.NewRecorder()
        req := httptest.NewRequest(http.MethodGet, "/_testStub", nil)
        c := e.NewContext(req, rec)

        

        req.Header.Set("Cookie", ACCESS_TOKEN_COOKIE_KEY + "=hello world")
        stubRequest := RequireAuthorization(handlerStub, false)

        assert.Error(t, stubRequest(c))
//         if assert.NOError(t, stubRequest(c)) {
//                 respBody := rec.Body.String()
//                 assert.Equal(t, respBody, HANDLER_STUB_SUCCESS)
//         }
 }

 // Provide a token for a REGULAR_USER and pass `true` in for `requireAdmin `

 func TestMiddlewareAdminToken(t *testing.T) {
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
        stubRequest := RequireAuthorization(handlerStub, true)

        assert.Error(t, stubRequest(c))

}