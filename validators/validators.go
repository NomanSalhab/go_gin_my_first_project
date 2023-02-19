package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateFullUserName(field validator.FieldLevel) bool {
	return len(field.Field().String()) > 8 /*strings.Contains(field.Field().String(), " ")*/
}

func ValidateUserIdNotZero(field validator.FieldLevel) bool {
	fmt.Println(field.Field().String())
	return field.Field().String() != "0"
}
