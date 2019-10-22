package error_utils

import "fmt"

type InvalidResponseError struct {
	code uint8
	err  string
}

func NewInvalidResponseError(code uint8, err string) *InvalidResponseError {
	return &InvalidResponseError{
		code: code,
		err:  err,
	}
}

func (err *InvalidResponseError) Error() string {
	return fmt.Sprintf("InvalidResponse: code=%d, err=%s", err.code, err.err)
}
