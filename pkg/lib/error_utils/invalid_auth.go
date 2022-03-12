package error_utils

import "fmt"

type InvalidAuthError struct {
	data interface{}
}

func NewInvalidAuthError(data interface{}) *InvalidAuthError {
	return &InvalidAuthError{
		data: data,
	}
}

func NewInvalidAuthErrorf(format string, args ...interface{}) *InvalidAuthError {
	return &InvalidAuthError{
		data: fmt.Sprintf(format, args...),
	}
}

func (err *InvalidAuthError) Error() string {
	return fmt.Sprintf("InvalidAuth: %v", err.data)
}
