package utils

import "errors"

var ErrorsMap = map[string]error{
	"userExistsError":    errors.New("user already exists"),
	"serverError":        errors.New("error in server, try again later"),
	"userNotFoundError":  errors.New("user not found"),
	"passwordWrongError": errors.New("password is wrong"),
}
