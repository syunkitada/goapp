package error_utils

import "fmt"

type TimeoutExceededError struct {
	data interface{}
}

func NewTimeoutExceededError(data interface{}) *TimeoutExceededError {
	return &TimeoutExceededError{
		data: data,
	}
}

func (err *TimeoutExceededError) Error() string {
	return fmt.Sprintf("TimeoutExceeded: %v", err.data)
}
