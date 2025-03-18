package api

import (
	"fmt"
	"net/http"
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

const defaultInternalErrorMsg = "Internal Error"

type Error struct {
	StatusCode int   `json:"-"`
	Message    any   `json:"message"`
	Err        error `json:"-"`
}

func (e Error) Error() string {
	return fmt.Sprintf("[API] %s\n%s", e.Message, e.Err)
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

func NewInternalError(err error, message ...any) *Error {
	if len(message) == 0 {
		message[0] = defaultInternalErrorMsg
	}
	return newError(http.StatusInternalServerError, err, message[0])
}
func NewNotFoundError(err error, message any) *Error {
	return newError(http.StatusNotFound, err, message)
}
func NewUnauthorizedError(err error, message any) *Error {
	return newError(http.StatusUnauthorized, err, message)
}
func NewBadRequestError(err error, message any) *Error {
	return newError(http.StatusBadRequest, err, message)
}
func NewDuplicateError(err error, message any) *Error {
	return newError(http.StatusConflict, err, message)
}
func NewForbiddenError(err error, message any) *Error {
	return newError(http.StatusForbidden, err, message)
}
