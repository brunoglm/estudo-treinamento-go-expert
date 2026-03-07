package errors

type Error struct {
	Message string
	Err     string
}

func (ie *Error) Error() string {
	return ie.Message
}

func NewNotFoundError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "not_found",
	}
}

func NewInternalServerError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "internal_server_error",
	}
}

func NewBadRequestError(message string) *Error {
	return &Error{
		Message: message,
		Err:     "bad_request",
	}
}
