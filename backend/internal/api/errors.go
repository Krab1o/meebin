package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Krab1o/meebin/internal/service"
)

type Error struct {
	StatusCode int   `json:"-"`
	Message    any   `json:"message"`
	Err        error `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("api error: %d", e.StatusCode)
}

func newError(statusCode int, message any, err error) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func NewNotFoundError(message any, err error) *Error {
	return newError(http.StatusNotFound, message, err)
}
func NewUnauthorizedError(message any, err error) *Error {
	return newError(http.StatusUnauthorized, message, err)
}
func NewInternalError(message any, err error) *Error {
	return newError(http.StatusInternalServerError, message, err)
}
func NewBadRequestError(message any, err error) *Error {
	return newError(http.StatusBadRequest, message, err)
}

func ErrorServiceToAPI(message any, err error) *Error {
	var serviceError *service.Error
	if errors.As(err, &serviceError) {
		switch serviceError.Type {
		case service.NotFound:
			return NewNotFoundError(message, err)
		case service.Semantic:
			return NewBadRequestError(message, err)
		case service.Unauthorized:
			return NewUnauthorizedError(message, err)
		case service.Internal:
		default:
			return NewInternalError(message, err)
		}
	}
	return NewInternalError(message, err)
}
