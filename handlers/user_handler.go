package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Harsh-apk/notesPostgres/db"
	"github.com/Harsh-apk/notesPostgres/types"
	"github.com/Harsh-apk/notesPostgres/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	db      db.UserDB
	envData *types.ENV_DATA
}

func NewUserHandler(db db.UserDB, data *types.ENV_DATA) UserHandler {
	return UserHandler{
		db:      db,
		envData: data,
	}
}

func (u *UserHandler) HandleGetAllUsers(c *fiber.Ctx) error {
	users, err := u.db.ReadAllUsers()
	if err != nil {
		c.Context().SetStatusCode(http.StatusNotFound)
		return c.JSON(map[string]string{"message": "no records found"})
	}
	c.Context().SetStatusCode(http.StatusOK)
	return c.JSON(map[string]([]types.User){"data": *users})
}

func (u *UserHandler) HandleAuthenticateEmail(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusBadRequest)
		return c.JSON(map[string]string{"message": "Something went wrong 1"})
	}
	err = u.db.AuthenticateUser(uint(id))
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusUnauthorized)
		return c.JSON(map[string]string{"message": err.Error()})
	}
	c.Context().SetStatusCode(fiber.StatusOK)
	return c.JSON(map[string]string{"message": "Successfully authorized"})

}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var user types.User
	err := c.BodyParser(&user)
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusBadRequest)
		return c.JSON(map[string]string{"message": "Something went wrong"})
	}
	err = u.db.CreateNewUser(&user)
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusUnprocessableEntity)
		return c.JSON(map[string]string{"message": err.Error()})
	}
	link := fmt.Sprintf("http://127.0.0.1:3000/auth/%d", user.ID)
	err = utils.SendSMS(&user.Email, u.envData, &link)
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusUnprocessableEntity)
		return c.JSON(map[string]string{"message": "Something went wrong while authenticating"})
	}
	c.Context().SetStatusCode(fiber.StatusOK)
	return c.JSON(map[string]string{"message": "Successfully created new account"})

}
