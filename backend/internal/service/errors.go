package service

import (
	"errors"
	"fmt"

	rep "github.com/Krab1o/meebin/internal/repository"
)

var (
	ErrNotFound    = errors.New("[SERVICE] Item not found")
	ErrUnautorized = errors.New("[SERVICE] Unauthorized")
	ErrInternal    = errors.New("[SERVICE] Internal Error")
	ErrDuplicate   = errors.New("[SERVICE] Item already exists")
	ErrForbidden   = errors.New("[SERVICE] Forbidden")
	ErrNoUpdate    = errors.New("[SERVICE] No fields to update")
)

func NewNotFoundError(err error, messages ...string) error {
	return fmt.Errorf("%w\n%v", ErrNotFound, err)
}
func NewUnauthorizedError(err error, messages ...string) error {
	return fmt.Errorf("%w\n%v", ErrUnautorized, err)
}
func NewInternalError(err error, messages ...string) error {
	return fmt.Errorf("%w\n%v", ErrInternal, err)
}
func NewDuplicateError(err error, messages ...string) error {
	return fmt.Errorf("%w\n%v", ErrDuplicate, err)
}
func NewForbiddenError(err error, messages ...string) error {
	return fmt.Errorf("%w\n%v", ErrForbidden, err)
}
func NewNoUpdateError(err error) error {
	return fmt.Errorf("%w\n%v", ErrNoUpdate, err)
}

func ErrorDBToService(
	err error,
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
