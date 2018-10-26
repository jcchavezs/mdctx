# mdctx

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