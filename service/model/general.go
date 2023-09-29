package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
