# notparam

[![pkg.go.dev][gopkg-badge]][gopkg]

`notparam` restricts not to use type parameters in declaration of  functions and types.

Type parameter has not been released yet. You can try it with gotip.

## Install

You can get `notparam` by `go install` command (Go 1.18 and higher).

```bash
$ go install github.com/gostaticanalysis/notparam/cmd/notparam@latest
```

## How to use

`notparam` run with `go vet` as below when Go is 1.18 and higher.

```bash
$ go vet -vettool=$(which notparam) ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [notparam.Analyzer](https://pkg.go.dev/github.com/gostaticanalysis/notparam/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/notparam
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/notparam?status.svg
