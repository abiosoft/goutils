# goutils

A set of personal utilities for Go.

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abiosoft/goutils)
[![Build Status](https://drone.io/github.com/abiosoft/goutils/status.png)](https://drone.io/github.com/abiosoft/goutils/latest)
[![Coverage Status](https://coveralls.io/repos/abiosoft/goutils/badge.svg)](https://coveralls.io/r/abiosoft/goutils)

I will keep adding more packages as time goes on and the need arises.

### Packages

#### 1. Set

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abiosoft/goutils/set)

##### Example
```go
import "github.com/abiosoft/goutils/set"

s := set.New()

// stores distinctly, repetitions have no effect.
s.AddAll(1, 1, 2, 2, 3, 6, 2, 1, "a", "b", "a")

s.Size() // 6

list := []interface{}{1, 2, 3, 6, "a", "b"}
s.ContainsList(list) // true

iter := s.Iterator()
for iter.HasNext() {
    // do something with iter.Value()
    ...
}
```

#### 2. Environment Variable

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/abiosoft/goutils/env)

##### Example
```go
import "github.com/abiosoft/goutils/env"

var vars env.EnvVar
vars.Set("GOPATH", "$HOME/go")
vars.Get("GOPATH") // $HOME/go
vars.String() // GOPATH=$HOME/go

vars.Set("GOOS", "darwin")
vars.Get("GOOS") // darwin
vars.String() // GOPATH=$HOME/go\nGOOS=darwin

// though env.EnvVar is a slice, adding an invalid string has no effect.
vars = append(vars, "SOME STRING")
vars.String() // GOPATH=$HOME/go\nGOOS=darwin
```
