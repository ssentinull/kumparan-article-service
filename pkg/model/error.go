package model

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest     = errors.New("the data type sent by the frontend is mismatched with the requirements in the backend")
	ErrInternalServer = errors.New("an error occured in the backend process")
)

func GetErrorStatusCode(err error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
