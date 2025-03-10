package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Krab1o/meebin/internal/repository"
)

type ErrorType int

type Error struct {
	Type    ErrorType
	Err     error
	Message string
}

const (
	Semantic ErrorType = iota
	Duplicate
	NotFound
	Unauthorized
	Internal
)

func (e ErrorType) String() string {
	switch e {
	case Semantic:
		return "Semantic Error"
	case Duplicate:
		return "Duplicate item"
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
	return fmt.Sprintf("Service error: %s", e.Message)
}

func (e Error) Unwrap() error {
	return e.Err
}

func newError(errType ErrorType, err error, messages ...string) *Error {
	var messageBuilder strings.Builder
	if len(messages) == 0 {
		messages = append(messages, errType.String())
	}
	for i, message := range messages {
		messageBuilder.WriteString(message)
		if i < len(messages)-1 {
			messageBuilder.WriteString("; ")
		}
	}

	return &Error{
		Type:    errType,
		Err:     err,
		Message: messageBuilder.String(),
	}
}

func NewSemanticError(err error, messages ...string) *Error {
	return newError(Semantic, err, messages...)
}
func NewNotFoundError(err error, messages ...string) *Error {
	return newError(NotFound, err, messages...)
}
func NewUnauthorizedError(err error, messages ...string) *Error {
	return newError(Unauthorized, err, messages...)
}
func NewInternalError(err error, messages ...string) *Error {
	return newError(Internal, err, messages...)
}
func NewDuplicateError(err error, messages ...string) *Error {
	return newError(Duplicate, err, messages...)
}

func defaultAction(dbError *repository.Error) *Error {
	switch dbError.Type {
	case repository.NotFound:
		return NewNotFoundError(dbError)
	case repository.Duplicate:
		return NewDuplicateError(dbError)
	default:
		return NewInternalError(dbError)
	}
}

func ErrorDBToService(
	err error,
	customAction func(*repository.Error) *Error,
) *Error {
	var dbError *repository.Error
	if errors.As(err, &dbError) {
		if customAction != nil {
			return customAction(dbError)
		} else {
			return defaultAction(dbError)
		}
	}
	return NewInternalError(err)
}
