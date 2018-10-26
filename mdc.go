package mdctx

import (
	"context"
)

type key string
type keyVals []interface{}

var mdcKey key = "mdc"

type mdcLogger struct {
	ctx    context.Context
	logger Logger
}

func (l *mdcLogger) Log(kvs ...interface{}) error {
	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}
	if ctxKvs := get(l.ctx); len(ctxKvs) != 0 {
		kvs = append(ctxKvs, kvs...)
	}
	return l.logger.Log(kvs...)
}

func get(ctx context.Context) keyVals {
	rawVal := ctx.Value(mdcKey)
	val, _ := rawVal.(keyVals)
	return val
}

// Add appends values into the context and they can be retrieved
// from the logger returned by With
func Add(ctx context.Context, kvs ...interface{}) context.Context {
	if len(kvs) == 0 {
		return ctx
	}

	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}

	kvs = append(get(ctx), kvs...)
	return context.WithValue(ctx, mdcKey, keyVals(kvs))
}

// Clear clears all MDC elements into the context
func Clear(ctx context.Context) context.Context {
	return context.WithValue(ctx, mdcKey, keyVals{})
}

// With returns a new contextual logger with keyvals prepended to those passed
// to calls to Log. If logger is also a contextual logger created by With or
// keyvals is appended to the existing context.
//
// The returned Logger replaces all value elements (odd indexes) containing a
// Valuer with their generated value for each call to its Log method.
func With(ctx context.Context, logger Logger) Logger {
	return &mdcLogger{ctx: ctx, logger: logger}
}
