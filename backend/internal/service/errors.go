package service

import (
	"errors"
	"fmt"

	rep "github.com/Krab1o/meebin/internal/repository"
)

var (
	ErrNotFound    = errors.New("[SERVICE] Item not found\n%w")
	ErrUnautorized = errors.New("[SERVICE] Unauthorized\n%w")
	ErrInternal    = errors.New("[SERVICE] Internal Error\n%w")
	ErrDuplicate   = errors.New("[SERVICE] Item already exists\n%w")
	ErrForbidden   = errors.New("[SERVICE] Forbidden\n%w")
	ErrNoUpdate    = errors.New("[SERVICE] No fields to update\n%w")
)

func NewNotFoundError(err error, messages ...string) error {
	return fmt.Errorf(ErrNotFound.Error(), err)
}
func NewUnauthorizedError(err error, messages ...string) error {
	return fmt.Errorf(ErrUnautorized.Error(), err)
}
func NewInternalError(err error, messages ...string) error {
	return fmt.Errorf(ErrInternal.Error(), err)
}
func NewDuplicateError(err error, messages ...string) error {
	return fmt.Errorf(ErrDuplicate.Error(), err)
}
func NewForbiddenError(err error, messages ...string) error {
	return fmt.Errorf(ErrForbidden.Error(), err)
}
func NewNoUpdateError(err error) error {
	return fmt.Errorf(ErrNoUpdate.Error(), err)
}

func ErrorDBToService(
	err error,
	// customAction func(*repository.Error) *Error,
) error {
	switch {
	case errors.Is(err, rep.ErrNotFound):
		return NewNotFoundError(err)
	case errors.Is(err, rep.ErrDuplicate):
		return NewDuplicateError(err)
	case errors.Is(err, rep.ErrInternal):
		return NewInternalError(err)
	default:
		return NewInternalError(err)
	}
}
