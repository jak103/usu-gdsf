package api

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func user(c echo.Context) error {
	return c.JSON(http.StatusOK, "User get handler")
}

func register(c echo.Context) error {
	// User registration screen
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	//hash password
	p := hashParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	passwordHash, err := generateEncodedPassword(c.FormValue("password"), p)
	if err != nil {
		log.WithError(err).Error("Password not encoded")
		return err
	}

	//sanitize birthday
	//What format am I getting this data?
	birthday, err := sanitizeBirthdayInput(c.FormValue("birthday"))
	if err != nil {
		log.WithError(err).Error("Invalid birthday entry")
		return err
	}

	newUser := models.User{
		c.FormValue("email"),
		passwordHash,
		c.FormValue("firstName"),
		c.FormValue("lastName"),
		birthday,
		time.Now(),
	}

	db.CreateUser(newUser)

	//Generate authentication tokens

	return c.JSON(http.StatusOK, "User registration handler")
}

// https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go
func generateEncodedPassword(password string, p hashParams) (encodedHash string, err error) {

	salt, err := generateSalt(p.saltLength)
	if err != nil {
		log.Error("Salt was not created")
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generateSalt(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Error("Salt was not created")
		return nil, err
	}

	return b, nil
}

//Added this function in case the I need to change how birthday sanitation is done
func sanitizeBirthdayInput(input string) (time.Time, error) {
	birthday, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return time.Now(), err
	}

	return birthday, nil
}

func verifyPassword(hashedString string, passwordInput string) bool {

	p, salt, hash, err := decodeHash(hashedString)
	if err != nil {
		return false
	}

	hashedPassword := argon2.IDKey([]byte(passwordInput), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	var notEqual byte = 0

	if len(hash) == 0 || len(hashedPassword) == 0 {
		return false
	}

	for i := 0; i < len(hash) && i < len(hashedPassword); i++ {
		notEqual |= hash[i] ^ hashedPassword[i]
	}

	return notEqual == 0 && len(hash) == len(hashedPassword)
}

func decodeHash(encodedHash string) (p *hashParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid Hash")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible version")
	}

	p = &hashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func downloads(c echo.Context) error {
	// Return games that a user has downloaded/played
	return c.JSON(http.StatusOK, "User downloads handler")
}

func init() {
	log.Info("Running user init")

	registerRoute(route{method: http.MethodGet, path: "/user", handler: user})
	registerRoute(route{method: http.MethodGet, path: "/user/register", handler: register})
	registerRoute(route{
		method:      http.MethodGet,
		path:        "user/downloads",
		handler:     downloads,
		requireAuth: true,
	})
}
