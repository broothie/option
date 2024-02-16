# `option`

[![Go Reference](https://pkg.go.dev/badge/github.com/broothie/option.svg)](https://pkg.go.dev/github.com/broothie/option)
[![Go Report Card](https://goreportcard.com/badge/github.com/broothie/option)](https://goreportcard.com/report/github.com/broothie/option)
[![codecov](https://codecov.io/gh/broothie/option/graph/badge.svg?token=dud3HQaZbN)](https://codecov.io/gh/broothie/option)
[![gosec](https://github.com/broothie/option/actions/workflows/gosec.yml/badge.svg)](https://github.com/broothie/option/actions/workflows/gosec.yml)

This package aims to make it easy to use the [options pattern](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) in Go.
It uses generics to ensure type safety, and has no dependencies outside the standard library.

## Getting Started

### Go Get

```bash
go get github.com/broothie/option
```

### Import

```go
import "github.com/broothie/option"
```

## Documentation

The API is small, so taking a look at the [package index on pkg.go.dev](https://pkg.go.dev/github.com/broothie/option#pkg-index) should get you up to speed.

## Usage

Let's say you have a `Server` that you'd like to be configurable:

```go
package server

import (
	"database/sql"
	"log/slog"
)

type Server struct {
	logger *slog.Logger
	db     *sql.DB
}
```

First, provide callers with option builders:

```go
func Logger(logger *slog.Logger) option.Func[*Server] {
	return func(server *Server) (*Server, error) {
		server.logger = logger
		return server, nil
	}
}

func DB(name string) option.Func[*Server] {
	return func(server *Server) (*Server, error) {
		db, err := sql.Open("pg", name)
		if err != nil {
			return nil, err
		}
		
		server.db = db
		return server, nil
	}
}
```

Then define a constructor that accepts a variadic argument of type `option.Option[*Server]`.
In it, use `option.Apply` to run the new server instance through the provided options.

```go
func New(options ...option.Option[*Server]) (*Server, error) {
	return option.Apply(new(Server), options...)
}
```

Now, callers can use the options pattern when instantiating your server:

```go
srv, err := server.New(
	server.Logger(slog.New(slog.NewTextHandler(os.Stdout, nil))),
	server.DB("some-connection-string"),
)
```

### Implementing `Option`

Let's say you want your server to always respond with a set of HTTP headers:

```go
type Server struct {
	headers http.Header
}
```

You can create a custom `Option` for configuring this headers like this:

```go
type Headers http.Header

func (h Headers) Apply(server *Server) (*Server, error) {
	server.headers = http.Header(h)
	return server, nil
}
```

Now, a caller can configure their server's headers like this:

```go
srv, err := server.New(server.Headers{
	"Content-Type": {"application/json"},
	"X-From": {"my-app"}
})
```
