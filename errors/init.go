package errors

import (
	"errors"
)

type IError interface {
	SetErrors(errors map[string]*ErrorItem)
}

type ErrorItem struct {
	Key     string                `json:"key,omitempty"`
	Title   string                `json:"title,omitempty"`
	Message string                `json:"message,omitempty"`
	Errors  map[string]*ErrorItem `json:"errors,omitempty"`
}

type Error struct {
	error
	IError `json:"-"`

	Status  int                   `json:"-"`
	Message string                `json:"message,omitempty"`
	Errors  map[string]*ErrorItem `json:"errors,omitempty"`
}

const (
	HttpOK            = 200
	BadRequestError   = 400
	NotFoundError     = 404
	UnauthorizedError = 401
	ForbiddenError    = 403
	InternalError     = 500
)

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) SetErrors(errors map[string]*ErrorItem) {
	e.Errors = errors
}

func IsStatus(err error, status int) bool {
	if e, ok := err.(*Error); ok {
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
