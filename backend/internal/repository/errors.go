package repository

import (
	"errors"
	"fmt"
)

// SQL error codes
const (
	SQLCodeDuplicate = "23505"
)

var (
	ErrNotFound  = errors.New("[DB] Item not found")
	ErrInternal  = errors.New("[DB] Internal error")
	ErrDuplicate = errors.New("[DB] Item already exists")
)

func NewNotFoundError(err error) error {
	return fmt.Errorf("%w\n[SOURCE] %v", ErrNotFound, err)
}
func NewInternalError(err error) error {
	return fmt.Errorf("%w\n[SOURCE] %v", ErrInternal, err)
}
func NewDuplicateError(err error) error {
	return fmt.Errorf("%w\n[SOURCE] %v", ErrDuplicate, err)
}
