package utils

import "errors"

var ErrorsMap = map[string]error{
	"bookNotFoundError":   errors.New("book not found"),
	"serverError":         errors.New("error in server, try again later"),
	"authorNotFoundError": errors.New("author not found"),
}
