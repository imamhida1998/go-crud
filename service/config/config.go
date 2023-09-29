package config

import (
	"go-crud/service/model"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
}

func (c *Config) InitEnv() error {
	err := godotenv.Load("crud.env")
	if err != nil {
		return err
	}
	return err
}

func (c *Config) CatchError(err error) {
	if err != nil {
		panic(any(err))
	}
}

func (c *Config) GetDBConfig() model.DBConfig {
	return model.DBConfig{
		DBName:   os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
}
