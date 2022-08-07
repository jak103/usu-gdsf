package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/log"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
)

func User(c echo.Context) error {
	user, _ := VerifyUser(c)

	if user != nil {
		GenerateTokenAndSetCookie(user, c)
		return nil
	}

	log.Error("Unable to login user %s.", user.Displayname)
	return c.JSON(http.StatusUnauthorized, "Invalid user "+user.Displayname)
}

func GetUserByID(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}
	userID, _ := uuid.Parse(c.Param("ID"))

	if result, err := db.GetUserByID(userID); err != nil {
		log.Error("An error occurred while getting game records: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func CreateUser(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()

	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	u := *new(models.User)
	if err = c.Bind(&u); err != nil {
		return err
	}
	u.SetUUID()

	if err := db.CreateUser(u); err != nil {
		log.Error("An error occurred while creating user records %v", err)
		return err
	} else {
		GenerateTokenAndSetCookie(&u, c)
		return c.JSON(http.StatusOK, "New User Added")
	}
}

func ReturnAllUsers(c echo.Context) error {
	db, err := db.NewDatabaseFromEnv()
	if err != nil {
		log.WithError(err).Error("Unable to use database")
		return err
	}

	if result, err := db.GetAllUsers(); err != nil {
		log.Error("An error occurred while getting a record of all users: %v", err)
		return err
	} else {
		return c.JSON(http.StatusOK, []interface{}{result})
	}
}

func init() {
	log.Info("Running user init")

	registerRoute(route{method: http.MethodGet, path: "/user/:id", handler: GetUserByID})
	registerRoute(route{method: http.MethodPost, path: "/login", handler: User})
	registerRoute(route{method: http.MethodPost, path: "/register", handler: CreateUser})
	registerRoute(route{method: http.MethodGet, path: "/user", handler: ReturnAllUsers})
}
