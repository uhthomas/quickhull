# quickhull
[![GoDoc](https://godoc.org/github.com/uhthomas/quickhull?status.svg)](https://godoc.org/github.com/uhthomas/quickhull)

A fun integer-bound quickhull implementation.

## Usage
```go
S := []Point{{1, 2}, {3, 4}}
H := quickhull.Find(S)
```