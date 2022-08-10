package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	uuid "github.com/google/uuid"
	models "github.com/jak103/usu-gdsf/models"
	echo "github.com/labstack/echo/v4"
	assert "github.com/stretchr/testify/assert"
)

func TestGetKey(t *testing.T) {
	os.Setenv("USUGDSF_AUTH_TOKEN", "fakeHwVJzEu2wTRBK4wNu80C")
	defer os.Unsetenv("USUGDSF_AUTH_TOKEN")

	jwtkey := GetKey()

	assert.Equal(t, jwtkey, "fakeHwVJzEu2wTRBK4wNu80C")
}

func TestVerifyUser(t *testing.T) {
	// os.Setenv("DB_TYPE", "mock")
	// defer os.Unsetenv("DB_TYPE")

	// os.Setenv("USUGDSF_AUTH_TOKEN", "fakeHwVJzEu2wTRBK4wNu80C")
	// defer os.Unsetenv("USUGDSF_AUTH_TOKEN")

	// f := make(url.Values)
	// f.Set("username", "Default")
	// f.Set("password", "default")

	// e := echo.New()

	// req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	// req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationForm)

	// c := e.NewContext(req, nil)

	// validUser, _ := VerifyUser(c)

	// assert.NotNil(t, validUser)
}

func TestGenerateTokenAndSetCookie(t *testing.T) {
	e := echo.New()
	r := e.Router()

	r.Add(http.MethodGet, "/:testToken", func(echo.Context) error { return nil })

	req := httptest.NewRequest(http.MethodGet, "/:testToken", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	user := models.User{
		Username:    "Billy1234",
		Password:    "Bobby4321",
		Displayname: "BillyBob",
		Role:        models.Admin,
	}

	user.ID = uuid.New()

	assert.NotPanics(t, func() {
		GenerateTokenAndSetCookie(&user, c)
	})
}

func TestGenerateToken(t *testing.T) {
	os.Setenv("USUGDSF_AUTH_TOKEN", "fakeHwVJzEu2wTRBK4wNu80C")
	defer os.Unsetenv("USUGDSF_AUTH_TOKEN")

	e := echo.New()
	r := e.Router()

	r.Add(http.MethodGet, "/:testToken", func(echo.Context) error { return nil })

	req := httptest.NewRequest(http.MethodGet, "/:testToken", nil)
	rec := httptest.NewRecorder()

	e.NewContext(req, rec)

	user := models.User{
		Username:    "Billy1234",
		Password:    "Bobby4321",
		Displayname: "BillyBob",
		Role:        models.Admin,
	}

	user.ID = uuid.New()

	expires := time.Now().Add(time.Minute)
	token := GenerateToken(&user, expires)

	assert.NotEmpty(t, token)
}

func TestSetTokenCookie(t *testing.T) {
	os.Setenv("USUGDSF_AUTH_TOKEN", "fakeHwVJzEu2wTRBK4wNu80C")
	defer os.Unsetenv("USUGDSF_AUTH_TOKEN")

	e := echo.New()
	r := e.Router()

	r.Add(http.MethodGet, "/:testCookie", func(echo.Context) error { return nil })

	req := httptest.NewRequest(http.MethodGet, "/:foo", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	user := models.User{
		Username:    "Billy1234",
		Password:    "Bobby4321",
		Displayname: "BillyBob",
		Role:        models.Admin,
	}

	user.ID = uuid.New()

	expires := time.Now().Add(time.Minute)
	token := GenerateToken(&user, expires)

	SetTokenCookie(user.Displayname, token, expires, c)

	assert.Contains(t, rec.Header().Get(echo.HeaderSetCookie), user.Displayname)
}
