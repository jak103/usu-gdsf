package api

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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
	return os.Getenv("USUGDSF_AUTH_TOKEN")
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

	user, _ := db.GetUserByUserName(username)

	if user != nil {
		if user.Password == SecurePassword(password) {
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

	tokenString, err := token.SignedString([]byte(GetKey()))

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

func SecurePassword(password string) string {
	block, err := aes.NewCipher([]byte(GetKey()))

	if err != nil {
		log.Error("Unable to secure password")
		panic(err)
	}

	iv := []byte{25, 44, 66, 78, 54, 19, 87, 07, 24, 63, 35, 05, 68, 72, 93, 22}

	plainPassword := []byte(password)
	cfb := cipher.NewCFBEncrypter(block, iv)

	cipherText := make([]byte, len(plainPassword))

	cfb.XORKeyStream(cipherText, plainPassword)

	return base64.StdEncoding.EncodeToString(cipherText)
}
