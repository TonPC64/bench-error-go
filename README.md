# Bench Error golang

> just benchmark simple package


## Run Bench

```sh
go test -benchmem -bench=.
```

## new error with fmt.Errorf vs errors.New() vs pkg/errors.New()

```text
goos: darwin
goarch: amd64
pkg: github.com/TonPC64/benchgo
BenchmarkNewErrWithFmt-8               	30000000	        50.5 ns/op	       5 B/op	       1 allocs/op
BenchmarkNewErrWithErrors-8            	2000000000	         0.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewErrWithPKGError-8          	 3000000	       655 ns/op	     288 B/op	       2 allocs/op
BenchmarkNewErrWithFmtConcat-8         	10000000	       179 ns/op	      16 B/op	       1 allocs/op
BenchmarkNewErrWithErrorsConcat-8      	100000000	        19.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewErrWithPKGErrorsConcat-8   	 2000000	       764 ns/op	     352 B/op	       4 allocs/op
BenchmarkStackTraceErrorConCat-8       	20000000	       103 ns/op	      96 B/op	       3 allocs/op
BenchmarkStackTraceErrorConCatFmt-8    	 5000000	       294 ns/op	      96 B/op	       3 allocs/op
PASS
ok  	github.com/TonPC64/benchgo	14.784s
```
