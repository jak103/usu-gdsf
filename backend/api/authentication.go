package api

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
	db "github.com/jak103/usu-gdsf/db"
	log "github.com/jak103/usu-gdsf/log"
	models "github.com/jak103/usu-gdsf/models"
	echo "github.com/labstack/echo/v4"
)

func GetKey() string {
	return os.Getenv("USUGDSF_AUTH_KEY")
}

type Claims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

func VerifyUser(c echo.Context) (*models.User, error) {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to access database")
		return nil, err
	}
	db.Connect()
	defer db.Disconnect()

	// TODO make sure that this function is working once fully implmented
	user, _ := db.GetUserByUserName(username)

	if user != nil {
		if user.Password == password {
			return user, nil
		}
	}

	log.Error("Failed to verify user %s", username)
	return nil, err
}

func GenerateTokenAndSetCookie(user *models.User, c echo.Context) error {
	expirationTime := time.Now().Add(3 * time.Hour)

	token := GenerateToken(user, expirationTime)

	SetTokenCookie(user.Displayname, token, expirationTime, c)

	return c.JSON(http.StatusOK, "GeneratedTokensandSetCookies")
}

func GenerateToken(user *models.User, expirationTime time.Time) string {
	claims := &Claims{
		Name: user.Displayname,
		ID:   user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(GetKey())

	if err != nil {
		return "Issue generating token."
	}

	if tokenString != "" {
		return tokenString
	}

	log.Error("Unable to generate token for %s", user.Displayname)
	return ""
}

func SetTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	log.Info("Cookie generated")
	c.SetCookie(cookie)
}
