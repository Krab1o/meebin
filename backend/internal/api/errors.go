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

func newError(statusCode int, err error, message any) *Error {

	return &Error{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func NewNotFoundError(err error, message any) *Error {
	return newError(http.StatusNotFound, err, message)
}
func NewUnauthorizedError(err error, message any) *Error {
	return newError(http.StatusUnauthorized, err, message)
}
func NewInternalError(err error, message any) *Error {
	return newError(http.StatusInternalServerError, err, message)
}
func NewBadRequestError(err error, message any) *Error {
	return newError(http.StatusBadRequest, err, message)
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

func defaultAction(serviceError *service.Error) *Error {
	switch serviceError.Type {
	case service.NotFound:
		return NewNotFoundError(serviceError, serviceError.Message)
	case service.Semantic:
		return NewBadRequestError(serviceError, serviceError.Message)
	case service.Unauthorized:
		return NewUnauthorizedError(serviceError, serviceError.Message)
	default:
		return NewInternalError(serviceError, serviceError.Message)
	}
}

func ErrorServiceToAPI(
	err error,
	customAction func(*service.Error) *Error,
) *Error {
	var serviceError *service.Error
	if errors.As(err, &serviceError) {
		if customAction != nil {
			return customAction(serviceError)
		} else {
			return defaultAction(serviceError)
		}
	}
	return NewInternalError(err, "Internal Server Error")
}
