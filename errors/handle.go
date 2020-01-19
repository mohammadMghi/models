package errors

import (
	"fmt"
	gm "github.com/go-ginger/models"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func HandleError(err error) error {
	return Error{
		Message: err.Error(),
	}
}

func GetError(request gm.IRequest, status int, err ...error) error {
	var msg string
	if status == 0 {
		status = 500
	}
	if err != nil && len(err) > 0 {
		msg = err[0].Error()
	} else {
		switch status {
		case NotFoundError:
			msg = request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "NotFoundError",
					Other: "Requested information not found",
				},
			})
		case UnauthorizedError:
			msg = request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "UnauthorizedError",
					Other: "You are not authorized to access this section",
				},
			})
		default:
			msg = request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "InvalidRequestWithCode",
					Other: "Invalid request, code {{.Code}}",
				},
				TemplateData: map[string]string{
					"Code": string(status),
				},
			})
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

func GetNotFoundError(request gm.IRequest, messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{
			request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "NotFoundError",
					Other: "requested information not found",
				},
			}),
		}
	}
	errs := getErrors(messages...)
	return GetError(request, NotFoundError, errs...)
}

func GetUnAuthorizedError(request gm.IRequest, messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{
			request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "UnauthorizedError",
					Other: "You are not authorized to access this section",
				},
			}),
		}
	}
	errs := getErrors(messages...)
	return GetError(request, UnauthorizedError, errs...)
}

func GetForbiddenError(request gm.IRequest, messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{
			request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "ForbiddenError",
					Other: "Access to this section is denied",
				},
			}),
		}
	}
	errs := getErrors(messages...)
	return GetError(request, ForbiddenError, errs...)
}

// GetValidationError returns error associated with HTTP Validation errors
func GetValidationError(request gm.IRequest, messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{
			request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "ValidationError",
					Other: "Invalid request information",
				},
			}),
		}
	}
	errs := getErrors(messages...)
	return GetError(request, BadRequestError, errs...)
}

func GetInternalServiceError(request gm.IRequest, messages ...string) error {
	if messages == nil || len(messages) == 0 {
		messages = []string{
			request.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:    "UnknownError",
					Other: "Unknown error",
				},
			}),
		}
	}
	errs := getErrors(messages...)
	return GetError(request, InternalError, errs...)
}
