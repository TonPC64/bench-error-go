# Benchgo

> just benchmark simple package

## new error with fmt.Errorf vs errors.New()

```text
goos: darwin
goarch: amd64
pkg: github.com/TonPC64/benchgo
BenchmarkNewErrWithFmt-8               	30000000	        48.0 ns/op
BenchmarkNewErrWithErrors-8            	2000000000	         0.24 ns/op
BenchmarkNewErrWithPKGError-8          	 3000000	       565 ns/op
BenchmarkNewErrWithFmtConcat-8         	10000000	       168 ns/op
BenchmarkNewErrWithErrorsConcat-8      	100000000	        18.5 ns/op
BenchmarkNewErrWithPKGErrorsConcat-8   	 2000000	       633 ns/op
PASS
ok  	github.com/TonPC64/benchgo	9.887s

```