package request

import "go-crud/service/model"

type Register struct {
	Firstname string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=12"`
	Email     string `json:"email" binding:"required,email"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdatedUsers struct {
	Id int `uri:"id" binding:"required"`
}

type DetailUsers struct {
	Id int `uri:"id" binding:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=12"`
	Email    string `json:"email" binding:"required,email"`
	Users    model.User
}
