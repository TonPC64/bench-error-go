package main

import (
	"errors"
	"fmt"
	"net/http"

	pkgError "github.com/pkg/errors"
	"golang.org/x/xerrors"

	httpErrors "github.com/TonPC64/benchgo/httpErrors"
	stackErrors "github.com/TonPC64/benchgo/stackErrors"
)

func main() {
	fmt.Println(bar())
}

func newErrWithFmt() error {
	return fmt.Errorf("error")
}

func newErrWithErrors() error {
	return errors.New("error")
}

func newErrWithPKGError() error {
	return pkgError.New("error")
}

func newErrWithXErrors() error {
	return xerrors.New("error")
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

func newErrWithXErrorConcat(err error) error {
	return xerrors.Errorf("error %+v", err)
}

// Stacktrace V1
func stackTraceErrorConCat() error {
	return redisClient()
}

func redis() error {
	return errors.New("[Redis]: connection erorr")
}

func redisClient() error {
	err := redis()
	msg := "[RedisClient]: Unable to connect " + err.Error()
	return errors.New(msg)
}

// Stacktrace V2
func stackTraceErrorConCatFmt() error {
	return redisClient2()
}

func redisClient2() error {
	err := redis()
	msg := fmt.Sprintf("[RedisClient]: Unable to connect %+v", err)
	return errors.New(msg)
}

// custom error
func newHttpError() error {
	return httpErrors.New("error", httpErrors.HTTPResponse{
		Code:     http.StatusInternalServerError,
		Response: "error message",
	})
}

func foo() error {
	httperr := httpErrors.New("http error", httpErrors.HTTPResponse{Response: http.StatusText(http.StatusInternalServerError), Code: http.StatusInternalServerError})
	return stackErrors.New("httpError", httperr)
}

func bar() error {
	return stackErrors.New("error2", foo())
}
