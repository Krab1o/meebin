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

// TODO: add messages on service layer.
// take messages from service layer and send them here for api error, so
// client will see reasonable errors. if no reasonable message provided (empty)
// there should be inserted basic message like "internal error" or "bad request"

// TODO: think of message as optional argument
// probably messages may be an optional argument so is there any. if there is,
// we will insert it into the api error message and send it to client

// TODO: make api errors independent from service errors codes. make logic
// of creating any api error not dependent to service code moving it to handler.
// in handler error should be asserted as service error and then dependent on
// its code it should be processed in each handler corresponding its requirements

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
