package error_utils

import "fmt"

type InvalidRequestError struct {
	data interface{}
}

func NewInvalidRequestError(data interface{}) *InvalidRequestError {
	return &InvalidRequestError{
		data: data,
	}
}

func (err *InvalidRequestError) Error() string {
	return fmt.Sprintf("InvalidRequest: %v", err.data)
}
