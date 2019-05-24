# mdctx

[![Build Status](https://travis-ci.com/jcchavezs/mdctx.svg?branch=master)](https://travis-ci.com/jcchavezs/mdctx)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcchavezs/mdctx)](https://goreportcard.com/report/github.com/jcchavezs/mdctx)
[![GoDoc](https://godoc.org/github.com/jcchavezs/mdctx?status.svg)](https://godoc.org/github.com/jcchavezs/mdctx)
[![Sourcegraph](https://sourcegraph.com/github.com/jcchavezs/mdctx/-/badge.svg)](https://sourcegraph.com/github.com/jcchavezs/mdctx?badge)


Mapped Diagnostic Context (MDC) for Go logging

The idea of Mapped Diagnostic Context is to provide a way to enrich log messages with pieces of information that could be not available in the scope where the logging actually occurs, but that can be indeed useful to better track the execution of the program.

## Usage

```go
func Middleware(next http.Handler) http.Handler {
  	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Header.Get(RequestIDHeader)
		if val != "" {
            ctx := mdctx.Add(r.Context(), "request_id", val)
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

## Providers

Providers allows you to include additional context to the logs coming from other
sources. E.g. if a middleware include the `request_id` in the context under a private
key, you can use the API of that middleware to inject that `request_id` in the mdc.

```go
package requestIDMiddleware

type reqIDKey string

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        val := r.Header.Get(RequestIDHeader)
        if val != "" {
            r.WithContext(context.WithValue(r.Context(), reqIDKey, val))
		}
		next.ServeHTTP(w, r)
	})
}
```

```go
package requestID

func Provider(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", ctx.Value(reqIDKey))
}
```

```go
package main

func main() {
	mdctx.RegisterProvider(requestID.Provider)
}
```