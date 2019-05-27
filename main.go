package main

import (
	"errors"
	"fmt"

	pkgError "github.com/pkg/errors"
)

func main() {}

func newErrWithFmt() error {
	return fmt.Errorf("error")
}

func newErrWithErrors() error {
	return errors.New("error")
}

func newErrWithPKGError() error {
	return pkgError.New("error")
}

func newErrWithFmtConcat(err error) error {
	return fmt.Errorf("error %v", err)
}

func newErrWithErrorsConcat(err error) error {
	return errors.New("error" + err.Error())
}

func newErrWithPKGErrorConcat(err error) error {
	return pkgError.Wrap(err, "error ")
}
