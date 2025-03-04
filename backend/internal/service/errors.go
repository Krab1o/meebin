package service

import (
	"errors"

	"github.com/Krab1o/meebin/internal/repository"
)

type ErrorType int

type Error struct {
	Type ErrorType
	Err  error
}

const (
	Semantic ErrorType = iota
	NotFound
	Unauthorized
	Internal
)

func (e ErrorType) String() string {
	switch e {
	case Semantic:
		return "Semantic Error"
	case NotFound:
		return "Not found"
	case Unauthorized:
		return "Unauthorized"
	case Internal:
		return "Internal Error"
	default:
		return "Unknown Error"
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func newError(errType ErrorType, err error) *Error {
	return &Error{
		Type: errType,
		Err:  err,
	}
}

func NewSemanticError(err error) *Error {
	return newError(Semantic, err)
}
func NewNotFoundError(err error) *Error {
	return newError(NotFound, err)
}
func NewUnauthorizedError(err error) *Error {
	return newError(Unauthorized, err)
}
func NewInternalError(err error) *Error {
	return newError(Internal, err)
}

func ErrorDBToService(err error) *Error {
	var dbError *repository.Error
	if errors.As(err, &dbError) {
		switch dbError.Type {
		case repository.NotFound:
			return NewNotFoundError(err)
		default:
			return NewInternalError(err)
		}
	}
	return NewInternalError(err)
}
