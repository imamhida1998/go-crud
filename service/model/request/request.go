package request

import "go-crud/service/model"

type Register struct {
	Firstname string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required,min=12"`
	Email     string `json:"email" validate:"required,email"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatedUsers struct {
	Id int `uri:"id" validate:"required"`
}

type DetailUsers struct {
	Id int `uri:"id" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=12"`
	Email    string `json:"email" validate:"required,email"`
	Users    model.User
}
