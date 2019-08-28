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
	HttpOK            = 200
	BadRequestError   = 400
	NotFoundError     = 404
	UnauthorizedError = 401
	ForbiddenError    = 403
	InternalError     = 500
)

func (e Error) Error() string {
	return e.Message
}

func getErrors(messages ...string) []error {
	errs := []error{}
	for _, message := range messages {
		errs = append(errs, errors.New(message))
	}
	return errs
}

func HandleError(err error) error {
	return Error{
		Message: err.Error(),
	}
}

func GetError(status int, err ...error) error {
	var msg string
	if status == 0 {
		status = 500
	}
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

func GetErrorFromInterface(err ...interface{}) error {
	var firstErr interface{}
	if err != nil && len(err) > 0 {
		firstErr = err[0]
	}
	if firstErr != nil {
		if e, ok := firstErr.(*Error); ok {
			return *e
		}
	}
	return Error{
		Status:  500,
		Message: fmt.Sprintf("%v", firstErr),
	}
}

func GetNotFoundError(messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{"Not found"}
	}
	errs := getErrors(messages...)
	return GetError(NotFoundError, errs...)
}

func GetUnAuthorizedError(messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{"You are not authorized to access this section"}
	}
	errs := getErrors(messages...)
	return GetError(UnauthorizedError, errs...)
}

func GetForbiddenError(messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{"Access to this section is denied"}
	}
	errs := getErrors(messages...)
	return GetError(ForbiddenError, errs...)
}

// GetValidationError returns error associated with HTTP Vlidation error
func GetValidationError(messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{"Your request is not valid"}
	}
	errs := getErrors(messages...)
	return GetError(BadRequestError, errs...)
}

func GetInternalServiceError(messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{"Unknown error"}
	}
	errs := getErrors(messages...)
	return GetError(InternalError, errs...)
}
