package error_utils

import (
	"fmt"
	"runtime"
	"strings"
)

type RecoveredError struct {
	data   interface{}
	stacks []string
}

func NewRecoveredError(data interface{}) *RecoveredError {
	var stacks []string
	for depth := 0; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		stacks = append(stacks, fmt.Sprintf("%d: %v:%d", depth, file, line))
	}

	return &RecoveredError{
		data:   data,
		stacks: stacks,
	}
}

func (err *RecoveredError) Error() string {
	return fmt.Sprintf("Recovered: %v\n%s", err.data, strings.Join(err.stacks, "\n"))
}
