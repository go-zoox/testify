# Testify - a simple test enhancement library

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/testify)](https://pkg.go.dev/github.com/go-zoox/testify)
[![Build Status](https://github.com/go-zoox/testify/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/testify/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/testify)](https://goreportcard.com/report/github.com/go-zoox/testify)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/testify/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/testify?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/testify.svg)](https://github.com/go-zoox/testify/issues)
[![Release](https://img.shields.io/github/testify/go-zoox/testify.svg?label=Release)](https://github.com/go-zoox/testify/testifys)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/testify
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/testify"
)

func TestSomething(t *testing.T) {
  // assert ok
  testify.Assert(t, true, "this should be true")

  // equality
  testify.Equal(t, 123, 123, "they should be equal")

  // inequality
  testify.NotEqual(t, 123, 456, "they should not be equal")

  // enums
  testify.Enums(t, []string{"foo", "bar", "baz", "qux"}, "quux", "quux should be in the list")

  // type
  testify.Type(t, 123, "int", "they should be the same type")

  // range
  testify.Range(t, 0, 123, 100, "100 should be in the range")
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).