package models

type Error struct {
	error

	Message string `json:"message,omitempty"`
}

func HandleError(err error) error  {
	return Error{
		Message: err.Error(),
	}
}
