package error_utils

import "fmt"

type NotFoundError struct {
	data interface{}
}

func NewNotFoundError(data interface{}) *NotFoundError {
	return &NotFoundError{
		data: data,
	}
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %v", err.data)
}
