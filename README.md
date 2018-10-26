# mdctx

[![Build Status](https://travis-ci.com/jcchavezs/mdctx.svg?branch=master)](https://travis-ci.com/jcchavezs/mdctx)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcchavezs/mdctx)](https://goreportcard.com/report/github.com/jcchavezs/mdctx)
[![GoDoc](https://godoc.org/github.com/jcchavezs/mdctx?status.svg)](https://godoc.org/github.com/jcchavezs/mdctx)
[![Sourcegraph](https://sourcegraph.com/github.com/jcchavezs/mdctx/-/badge.svg)](https://sourcegraph.com/github.com/jcchavezs/mdctx?badge)


Mapped Diagnostic Context for Go logging

## Usage

```go
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get(RequestIDHeader)
		if val != "" {
            mdctx.Add(r.Context(), "request_id", val)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
```

```go
func (r *repository) DoSomething(ctx context.Context, ...) {
    logger := mdctx.With(ctx, r.logger)
    ...
    logger.Log("key", "value")
}
```