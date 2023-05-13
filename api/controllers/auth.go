package controllers

import (
	"ffcs/api/utils"
	"ffcs/pkg/models"
	"ffcs/pkg/status"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var message status.StatusStruct
var Validate = validator.New() // validator

// signup user function

func (r Repository) SignupUser(c *fiber.Ctx) error {
	var signupuser models.SignupUser
	if err := c.BodyParser(&signupuser); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(message.Error(err.Error()))
	}

	if err := Validate.Struct(signupuser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(message.Error("Unable to validate the request"))
	}

	user := signupuser.GetUserDetials()
	if err := r.DB.Create(&user).Error; err != nil {
		if strings.HasPrefix("[NOS]", err.Error()) {
			return c.Status(http.StatusBadRequest).JSON(message.Error(fmt.Sprintf("No branch registered %s", user.Branch)))
		} else if strings.Contains(err.Error(), "duplicate key value violates") {
			return c.Status(http.StatusUnprocessableEntity).JSON(message.Error("The user already exists in the server"))
		} else {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(message.Error("internal server error occured"))
		}
	}
	key, err := r.Rdb.CreateSession(user.ID, user.IsAdmin, user.Branch)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(message.Error("Unable to create a session for the user"))
	}

	c.Cookie(&fiber.Cookie{
		Name:    "session_id",
		Value:   key,
		Expires: time.Now().Add(time.Hour),
	})

	return c.Status(http.StatusOK).JSON(message.Success("successfully created the user", fiber.Map{
		"user": user.Rollno,
	}))
}

// login user
func (r Repository) LoginUser(c *fiber.Ctx) error {
	var Logindata models.LoginUser

	if err := c.BodyParser(&Logindata); err != nil {

		return c.Status(http.StatusUnprocessableEntity).JSON(message.Error("Unable to procress the login data"))
	}
	var Userdata models.Users
	if err := r.DB.First(&Userdata, "rollno = ?", Logindata.Rollno).Error; err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(message.Error("unable to find the requested credentials"))
	}
	if Userdata.Password == utils.Hashes(Logindata.Password) {
		key, err := r.Rdb.CreateSession(Userdata.ID, Userdata.IsAdmin, Userdata.Branch)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(message.Error("Unable to create a session for the user"))
		}
		c.Cookie(&fiber.Cookie{
			Name:    "session_id",
			Value:   key,
			Expires: time.Now().Add(time.Hour),
		})

		return c.Status(http.StatusOK).JSON(message.Success("User authenticated successfull", nil))
	}
	return c.Status(http.StatusUnauthorized).JSON(message.Error("unable to verify the identity"))
}
