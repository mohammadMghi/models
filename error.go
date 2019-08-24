package models

import (
	"errors"
	"fmt"
)

type Error struct {
	error

	Status  int    `json:"-"`
	Message string `json:"message,omitempty"`
}

const (
	BadRequestError   = 404
	NotFoundError     = 404
	UnauthorizedError = 401
	ForbiddenError    = 403
	InternalError     = 500
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
		case NotFoundError:
			msg = "not found"
		case UnauthorizedError:
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

func GetUnAuthorizedError(err ...error) error {
	if err == nil || len(err) == 0 {
		err = []error{errors.New("You are not authorized to access this section")}
	}
	return GetError(UnauthorizedError, err...)
}

func GetForbiddenError(err ...error) error {
	if err == nil || len(err) == 0 {
		err = []error{errors.New("Access to this section is denied")}
	}
	return GetError(ForbiddenError, err...)
}

// GetValidationError returns error associated with HTTP Vlidation error
func GetValidationError(err ...error) error {
	if err == nil || len(err) == 0 {
		err = []error{errors.New("Your request is not valid")}
	}
	return GetError(BadRequestError, err...)
}

func GetInternalServiceError(err ...error) error {
	if err == nil || len(err) == 0 {
		err = []error{errors.New("Unknown error")}
	}
	return GetError(InternalError, err...)
}
