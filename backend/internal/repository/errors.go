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
	ErrNotFound  = errors.New("[DB] Item not found\n[SOURCE]%w")
	ErrInternal  = errors.New("[DB] Internal error\n[SOURCE]%w")
	ErrDuplicate = errors.New("[DB] Item already exists\n[SOURCE]%w")
)

func NewNotFoundError(err error) error {
	return fmt.Errorf(ErrNotFound.Error(), err)
}
func NewInternalError(err error) error {
	return fmt.Errorf(ErrNotFound.Error(), err)
}
func NewDuplicateError(err error) error {
	return fmt.Errorf(ErrDuplicate.Error(), err)
}
