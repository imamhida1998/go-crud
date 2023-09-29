package handler

import (
	"go-crud/service/helpers"
	"go-crud/service/model"
	"go-crud/service/model/request"
	"go-crud/service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	userUsecase usecase.UsercaseUser
}

func NewHandlerUser(user usecase.UsercaseUser) *HandlerUser {
	return &HandlerUser{user}
}

func (h *HandlerUser) RegistrationDataUser(c *gin.Context) {
	var input request.Register

	err := c.ShouldBindJSON(&input)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, MessageError)
		return
	}

	err = h.userUsecase.RegistrationUser(input)
	if err != nil {
		MessageError := gin.H{"errors": err}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}

	c.JSON(http.StatusOK, "succes")

}

func (h *HandlerUser) LoginDataUser(c *gin.Context) {
	var input request.Login
	input.Username = c.Query("username")
	input.Password = c.Query("password")

	Token, err := h.userUsecase.LoginUsers(input)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}

	Response := gin.H{"token": Token}

	c.JSON(http.StatusOK, Response)

}

func (h *HandlerUser) UpdateDataUser(c *gin.Context) {
	var (
		datausers request.UpdateUser
		input     request.UpdatedUsers
	)

	err := c.ShouldBindUri(&input)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}
	err = c.ShouldBindJSON(&datausers)
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusUnprocessableEntity, errorsMessage)
		return
	}

	currentUser := c.MustGet("CurrentUser").(model.User)
	datausers.Users = currentUser

	err = h.userUsecase.UpdateUsers(input.Id, datausers)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}

	c.JSON(http.StatusOK, "update sukses")

}

func (h *HandlerUser) DetailUsers(c *gin.Context) {
	var (
		input request.DetailUsers
	)

	err := c.ShouldBindUri(&input)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}

	currentUser := c.MustGet("CurrentUser").(model.User)
	Users := currentUser

	res, err := h.userUsecase.DetailUsers(input.Id, Users)
	if err != nil {
		MessageError := gin.H{"errors": err.Error()}
		c.JSON(http.StatusInternalServerError, MessageError)
		return
	}

	c.JSON(http.StatusOK, helpers.DetailUsersFormatter(res))

}
