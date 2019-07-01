package errors

import (
	"runtime"
)

type StackError struct {
	err string
}

func New(s string, es ...error) error {
	caller := getCallerFunctionName()
	err := StackError{err: "[" + caller + "]: " + s}

	for _, e := range es {
		if e != nil {
			err.err = err.err + " " + e.Error()
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
