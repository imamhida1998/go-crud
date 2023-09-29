package handler

import (
	"encoding/json"
	"go-crud/service/helpers"
	"go-crud/service/model"
	"go-crud/service/model/request"
	"go-crud/service/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type HandlerUser struct {
	session     *session.Session
	userUsecase usecase.UsercaseUser
}

func NewHandlerUser(user usecase.UsercaseUser, session *session.Session) *HandlerUser {
	return &HandlerUser{session, user}
}

func (h *HandlerUser) RegistrationDataUser(c *fiber.Ctx) error {
	var input request.Register

	err := c.BodyParser(&input)
	if err != nil {

		MessageError := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusBadRequest).JSON(MessageError)
		return err
	}
	err = helpers.ValidateStruct(input)
	if err != nil {

		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusUnprocessableEntity).JSON(errorsMessage)
		return err

	}

	err = h.userUsecase.RegistrationUser(input)
	if err != nil {
		MessageError := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(MessageError)
		return err
	}

	c.Status(http.StatusOK).JSON("succes")
	return nil
}

func (h *HandlerUser) LoginDataUser(c *fiber.Ctx) error {
	var input request.Login
	input.Username = c.Query("username")
	input.Password = c.Query("password")

	Token, err := h.userUsecase.LoginUsers(input)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)
		return err
	}
	Response := map[string]interface{}{
		"message": "login succes",
		"metadata": map[string]string{
			"token": Token,
		},
	}
	c.Status(http.StatusOK).JSON(Response)

	return nil
}

func (h *HandlerUser) UpdateDataUser(c *fiber.Ctx) error {
	var (
		datausers request.UpdateUser
		input     request.UpdatedUsers
	)

	Id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)
		return err
	}
	input.Id = Id

	err = c.BodyParser(&datausers)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusUnprocessableEntity).JSON(errorsMessage)
		return err
	}

	err = helpers.ValidateStruct(datausers)
	if err != nil {

		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusUnprocessableEntity).JSON(errorsMessage)
		return err

	}

	currentUser := h.session.Get("CurrentUser")

	bodyBytes, _ := json.Marshal(currentUser)
	err = json.Unmarshal(bodyBytes, &datausers.Users)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)

		return err
	}

	err = h.userUsecase.UpdateUsers(input.Id, datausers)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)

		return err
	}
	c.Status(http.StatusOK).JSON("update sukses")
	return nil
}

func (h *HandlerUser) DetailUsers(c *fiber.Ctx) error {
	var (
		input request.DetailUsers
		Users model.User
	)

	Id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)
		return err
	}
	input.Id = Id

	currentUser := h.session.Get("CurrentUser")
	log.Println(currentUser)

	bodyBytes, _ := json.Marshal(currentUser)
	err = json.Unmarshal(bodyBytes, &Users)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)

		return err
	}

	res, err := h.userUsecase.DetailUsers(input.Id, Users)
	if err != nil {
		errorsMessage := map[string]string{
			"errors": err.Error(),
		}
		c.Status(http.StatusInternalServerError).JSON(errorsMessage)
		return err
	}
	c.Status(http.StatusOK).JSON(res)

	return nil
}
