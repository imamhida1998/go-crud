package repository

import (
	"database/sql"
	"errors"
	"go-crud/service/model"
	"go-crud/service/model/request"
	"log"
)

type UserRepo interface {
	Registration(input request.Register) error
	Login(input request.Login) (res model.User, err error)
	GetUsersById(id int) (res model.User, err error)
	GetUsersByEmail(email string) (res model.User, err error)
	UpdateDataUsers(id int, data model.User) error
}

type repoUser struct {
	db *sql.DB
}

func NewRepoUser(db *sql.DB) *repoUser {
	return &repoUser{db}
}

func (r *repoUser) Registration(input request.Register) error {
	query := `
		insert 
			into 
		users 
			(name,
			username,
			password,
			email,
			created_at) 
		values 
			(?,
			 ?,
			 ?,
			 ?,
			NOW())
		`

	name := input.Firstname + " " + input.LastName

	_, err := r.db.Exec(query, name, input.Username, input.Password, input.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repoUser) Login(input request.Login) (res model.User, err error) {
	query := `select name,username,password,email,created_at from users where username = ? and password = ?`
	rows, err := r.db.Query(query, input.Username, input.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Name,
			&res.Username,
			&res.Password,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}

func (r *repoUser) GetUsersById(id int) (res model.User, err error) {
	query := `select name,username,password,email,created_at  from users where id = ?`
	rows, err := r.db.Query(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Name,
			&res.Username,
			&res.Password,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}

func (r *repoUser) UpdateDataUsers(id int, data model.User) error {
	query := `
		update
			users
		set
		name = ?,
		username = ?,
		password = ?,
		email = ?,
		updated_at = NOW()
	where
		id = ?
		`

	_, err := r.db.Exec(query, data.Name, data.Username, data.Password, data.Email, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *repoUser) GetUsersByEmail(email string) (res model.User, err error) {
	query := `select name,username,password,email,created_at  from users where email = ?`
	rows, err := r.db.Query(query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		errx := rows.Scan(
			&res.Name,
			&res.Username,
			&res.Password,
			&res.Email,
			&res.CreatedAt)
		if errx != nil {
			return res, errx
		}

	}
	return res, nil

}
