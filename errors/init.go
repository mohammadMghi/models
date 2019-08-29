package errors

import (
	"errors"
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

func IsStatus(err error, status int) bool {
	if e, ok := err.(Error); ok {
		return e.Status == status
	}
	return false
}

func getErrors(messages ...string) []error {
	errs := make([]error, 0)
	for _, message := range messages {
		errs = append(errs, errors.New(message))
	}
	return errs
}
