package error_utils

import "fmt"

type ConflictNotFoundError struct {
	data interface{}
}

func NewConflictNotFoundError(data interface{}) *ConflictNotFoundError {
	return &ConflictNotFoundError{
		data: data,
	}
}

func (err *ConflictNotFoundError) Error() string {
	return fmt.Sprintf("ConflictNotFound: %v should be exists, but not found", err.data)
}
