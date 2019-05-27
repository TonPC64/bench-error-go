package main

import (
	"errors"
	"fmt"
)

func main() {}

func newErrWithFmt() error {
	return fmt.Errorf("error")
}

func newErrWithErrors() error {
	return errors.New("error")
}

func newErrWithFmtConcat(err error) error {
	return fmt.Errorf("error %v", err)
}

func newErrWithErrorsConcat(err error) error {
	return errors.New("error" + err.Error())
}
