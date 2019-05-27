# Benchgo

> just benchmark simple package

## new error with fmt.Errorf vs errors.New()

```text
goos: darwin
goarch: amd64
pkg: github.com/TonPC64/benchgo
BenchmarkNewErrWithFmt-8            	30000000	        45.2 ns/op
BenchmarkNewErrWithErrors-8         	2000000000	         0.25 ns/op
BenchmarkNewErrWithFmtConcat-8      	10000000	       167 ns/op
BenchmarkNewErrWithErrorsConcat-8   	100000000	        17.4 ns/op
PASS
ok  	github.com/TonPC64/benchgo	5.547s
```