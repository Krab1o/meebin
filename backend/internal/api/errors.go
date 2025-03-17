package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Krab1o/meebin/internal/service"
)

// this are comments to make error work clearer

// messages also exist on service layer.
// take messages from service layer and send them here for api error, so
// client will see reasonable errors. if no reasonable message provided (empty)
// there should be inserted basic message like "internal error" or "bad request"

// message is an optional argument
// probably messages may be an optional argument so is there any. if there is,
// we will insert it into the api error message and send it to client

// api errors are independent from service errors codes.
// logic of creating any api error not dependent to service code (handler).
// in handler error should be asserted as service error and then dependent on
// its code it should be processed in each handler corresponding its requirements

// TODO: remove status codes from all error
// TODO: switch to errors.Is
type Error struct {
	StatusCode int   `json:"-"`
	Message    any   `json:"message"`
	Err        error `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("API error: %s", e.Message)
}

func (e Error) Unwrap() error {
	return e.Err
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
func NewDuplicateError(err error, message any) *Error {
	return newError(http.StatusUnprocessableEntity, err, message)
}

func defaultAction(serviceError *service.Error) *Error {
	switch serviceError.Type {
	case service.NotFound:
		return NewNotFoundError(serviceError, serviceError.Message)
	case service.Semantic:
		return NewBadRequestError(serviceError, serviceError.Message)
	case service.Unauthorized:
		return NewUnauthorizedError(serviceError, serviceError.Message)
	case service.Duplicate:
		return NewDuplicateError(serviceError, serviceError.Message)
	default:
		return NewInternalError(serviceError, serviceError.Message)
	}
}

// Returns API error depending on type of service error and sets proper message
// for it
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
