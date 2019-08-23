package models

import (
	"fmt"
)

type Error struct {
	error

	Status  int    `json:"-"`
	Message string `json:"message,omitempty"`
}

const (
	NOT_FOUND = 404
)

func (e Error) Error() string {
	return e.Message
}

func HandleError(err error) error {
	return Error{
		Message: err.Error(),
	}
}

func GetError(status int, err ...error) error {
	var msg string
	if err != nil && len(err) > 0 {
		msg = err[0].Error()
	} else {
		switch status {
		case NOT_FOUND:
			msg = "not found"
		default:
			msg = fmt.Sprintf("invalid request, code %d", status)
		}
	}
	return Error{
		Status:  status,
		Message: msg,
	}
}
