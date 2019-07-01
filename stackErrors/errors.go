package errors

import (
	"runtime"
)

type StackError struct {
	err string
}

func New(s string, e ...error) error {
	caller := getCallerFunctionName()
	err := StackError{err: "[" + caller + "]: " + s}

	if len(e) > 0 {
		if newErr, ok := e[0].(StackError); ok {
			err.err = err.err + " " + newErr.err
		}
	}

	return err
}

func (s StackError) Error() string {
	return s.err
}

func getCallerFunctionName() string {
	if pc, _, _, ok := runtime.Caller(2); ok {
		if details := runtime.FuncForPC(pc); details != nil {
			return details.Name()
		}
	}

	return "unknown"
}
