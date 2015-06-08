# goutils

A set of personal utilities for Go.

I will keep adding more packages as time goes on.

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abiosoft/goutils/set)

### Packages

#### 1. Set

[Documentation](https://godoc.org/github.com/abiosoft/goutils/set)

##### Example
```go
s := NewSet()

// stores distinctly, repetitions have no effect.
s.AddAll(1, 1, 2, 2, 3, 6, 2, 1, "a", "b", "a")

s.Size() // 6

list := []interface{}{1, 2, 3, 6, "a", "b"}
s.ContainsAll(list...) // true

iter := s.Iterator()
for iter.HasNext() {
    // do something with iter.Value()
    ...
}
```
