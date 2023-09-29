package db

import (
	"database/sql"
	"fmt"
	"go-crud/service/model"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectiontoMYSQL(config model.DBConfig) (*sql.DB, error) {
	Constring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	db, err := sql.Open("mysql", Constring)
	if err != nil {
		return nil, err
	}
	return db, nil
}
