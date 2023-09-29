package helpers

import (
	"go-crud/service/model"
	"go-crud/service/model/response"
)

func DetailUsersFormatter(data model.User) (res response.FormatterUsers) {
	res.Name = data.Name
	res.Username = data.Username
	res.Password = data.Password
	res.Email = data.Email
	return res

}
