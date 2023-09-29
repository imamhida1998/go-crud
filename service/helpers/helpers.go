package helpers

import (
	"fmt"

	"github.com/go-playground/validator"
)

func ValidateStruct(data interface{}) error {
	var validate = validator.New()

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
		return err
	}
	return err
}
