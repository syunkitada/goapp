package error_utils

import "fmt"

type RecoveredError struct {
	data interface{}
}

func NewRecoveredError(data interface{}) *RecoveredError {
	return &RecoveredError{
		data: data,
	}
}

func (err *RecoveredError) Error() string {
	return fmt.Sprintf("Recovered: %v", err.data)
}
