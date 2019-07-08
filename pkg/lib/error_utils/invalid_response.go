package error_utils

import "fmt"

type InvalidResponseError struct {
	data interface{}
}

func NewInvalidResponseError(data interface{}) *InvalidResponseError {
	return &InvalidResponseError{
		data: data,
	}
}

func (err *InvalidResponseError) Error() string {
	return fmt.Sprintf("InvalidResponse: %v", err.data)
}
