package models

type Error struct {
	error

	Message string `json:"message,omitempty"`
}

func (e Error) Error() string  {
	return e.Message
}

func HandleError(err error) error  {
	return Error{
		Message: err.Error(),
	}
}
