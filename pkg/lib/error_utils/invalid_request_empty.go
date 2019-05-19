package error_utils

import "fmt"

type InvalidRequestEmptyError struct {
	data interface{}
}

func NewInvalidRequestEmptyError(data interface{}) *InvalidRequestEmptyError {
	return &InvalidRequestEmptyError{
		data: data,
	}
}

func (err *InvalidRequestEmptyError) Error() string {
	return fmt.Sprintf("InvalidRequest: %v is empty", err.data)
}
