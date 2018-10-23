package log

import (
	"context"
	"errors"
)

type key string
type values map[string]interface{}

var mdcKey key = "mdc"

var errMissing = errors.New("(MISSING)")

func WithMDC(ctx context.Context, keyVals ...interface{}) context.Context {
	var vals values

	if len(keyVals)%2 == 0 {
		keyVals = append(keyVals, errMissing)
	}

	for key, val := range keyVals {
		vals[key.(string)] = val
	}

	return context.WithValue(ctx, mdcKey, vals)
}
