package main

import (
	"errors"
	"testing"
)

func BenchmarkNewErrWithFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newErrWithFmt()
	}
}

func BenchmarkNewErrWithErrors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newErrWithErrors()
	}
}

var err = errors.New("concat")

func BenchmarkNewErrWithFmtConcat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		newErrWithFmtConcat(err)
	}
}

func BenchmarkNewErrWithErrorsConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newErrWithErrorsConcat(err)
	}
}
