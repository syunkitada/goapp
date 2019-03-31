package error_utils

import "fmt"

type ConflictAlreadyExistsError struct {
	data interface{}
}

func NewConflictAlreadyExistsError(data interface{}) *ConflictAlreadyExistsError {
	return &ConflictAlreadyExistsError{
		data: data,
	}
}

func (err *ConflictAlreadyExistsError) Error() string {
	return fmt.Sprintf("ConflictAlreadyExists: %v", err.data)
}
