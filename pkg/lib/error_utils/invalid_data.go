package error_utils

import "fmt"

type InvalidDataError struct {
	key  string
	data interface{}
	msg  string
}

func NewInvalidDataError(key string, data interface{}, msg string) *InvalidDataError {
	return &InvalidDataError{
		key:  key,
		data: data,
		msg:  msg,
	}
}

func (err *InvalidDataError) Error() string {
	return fmt.Sprintf("InvalidData: %s %v: %s", err.key, err.data, err.msg)
}
