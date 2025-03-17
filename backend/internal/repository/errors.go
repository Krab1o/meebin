package repository

import "fmt"

type ErrorType int

type Error struct {
	Type    ErrorType
	Err     error
	Message string
}

// SQL error codes
const (
	SQLCodeDuplicate = "23505"
)

const (
	NotFound ErrorType = iota
	Internal
	Duplicate
)

func (e ErrorType) String() string {
	switch e {
	case NotFound:
		return "Not found"
	case Internal:
		return "Internal Error"
	case Duplicate:
		return "Duplicate Item"
	default:
		return "Unknown Error"
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("DB error: %s", e.Err.Error())
}

func (e Error) Unwrap() error {
	return nil
}

func newError(errType ErrorType, err error) *Error {
	return &Error{
		Type: errType,
		Err:  err,
	}
}

func NewNotFoundError(err error) *Error {
	return newError(NotFound, err)
}
func NewInternalError(err error) *Error {
	return newError(Internal, err)
}
func NewDuplicateError(err error) *Error {
	return newError(Duplicate, err)
}
