package usecase

import (
	"errors"
	"go-crud/service/model"
	"go-crud/service/model/request"
	"go-crud/service/repo"
)

type UsercaseUser interface {
	RegistrationUser(input request.Register) error
	LoginUsers(input request.Login) (string, error)
	UpdateUsers(id int, input request.UpdateUser) error
	DetailUsers(id int, account model.User) (res model.User, err error)
}

type usecaseUser struct {
	repo repo.UserRepo
	auth Auth
}

func NewUsecaseUser(repo repo.UserRepo, a Auth) *usecaseUser {
	return &usecaseUser{repo, a}
}

func (u *usecaseUser) RegistrationUser(input request.Register) error {
	err := u.repo.Registration(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *usecaseUser) LoginUsers(input request.Login) (string, error) {
	res, err := u.repo.Login(input)
	if err != nil {
		return "", err
	}
	if res.Username == "" {
		return "", errors.New("login gagal")
	}
	Token, err := u.auth.GenerateTokenJWT(res.Email)
	if err != nil {
		return "", err
	}

	return Token, nil
}

func (u *usecaseUser) UpdateUsers(id int, input request.UpdateUser) error {
	res, err := u.repo.GetUsersById(id)
	if err != nil {
		return err
	}
	if res.Username != input.Users.Username {
		return errors.New("login gagal")
	}
	res.Name = input.Name
	res.Email = input.Email
	res.Password = input.Password
	res.Username = input.Username
	err = u.repo.UpdateDataUsers(id, res)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecaseUser) DetailUsers(id int, account model.User) (res model.User, err error) {
	res, err = u.repo.GetUsersById(id)
	if err != nil {
		return res, err
	}
	if res.Username != account.Username {
		return res, errors.New("login gagal")
	}

	err = u.repo.UpdateDataUsers(id, res)
	if err != nil {
		return res, err
	}

	return res, nil
}
