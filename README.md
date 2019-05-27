# Benchgo

> just benchmark simple package


## Run Bench

```sh
go test -benchmem -bench=.
```

## new error with fmt.Errorf vs errors.New()

```text
goos: darwin
goarch: amd64
pkg: github.com/TonPC64/benchgo
BenchmarkNewErrWithFmt-8               	30000000	        48.9 ns/op	       5 B/op	       1 allocs/op
BenchmarkNewErrWithErrors-8            	2000000000	         0.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewErrWithPKGError-8          	 3000000	       524 ns/op	     288 B/op	       2 allocs/op
BenchmarkNewErrWithFmtConcat-8         	10000000	       158 ns/op	      16 B/op	       1 allocs/op
BenchmarkNewErrWithErrorsConcat-8      	100000000	        16.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewErrWithPKGErrorsConcat-8   	 2000000	       606 ns/op	     352 B/op	       4 allocs/op
PASS
ok  	github.com/TonPC64/benchgo	9.420s
```