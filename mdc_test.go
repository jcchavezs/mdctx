package mdctx

import (
	"context"
	"testing"
)

func TestAddMDCForEmptyContext(t *testing.T) {
	kvs := get(context.Background())
	if want, have := 0, len(kvs); want != have {
		t.Fatalf("unexpected keyVals length, want: %d, have: %d", want, have)
	}
}
func TestAddMDCSuccess(t *testing.T) {
	ctx := Add(context.Background(), "key", "val")
	kvs := get(ctx)
	if want, have := "val", kvs["key"]; want != have {
		t.Fatalf("unexpected key, want: %q, have: %q", want, have)
	}
}

type inspectLogger struct {
	kvs keyVals
}

func (l *inspectLogger) Log(kvs ...interface{}) error {
	if len(kvs)%2 != 0 {
		kvs = append(kvs, ErrMissingValue)
	}

	if l.kvs == nil {
		l.kvs = keyVals{}
	}

	for i := 0; i < len(kvs)/2; i++ {
		idx, ok := kvs[2*i].(string)
		if ok {
			l.kvs[idx] = kvs[2*i+1]
		}
	}
	return nil
}

func TestWithSuccessWhenHavingValues(t *testing.T) {
	ctx := Add(context.Background(), "key", "val")
	logger := &inspectLogger{}

	wLogger := With(ctx, logger)
	wLogger.Log("key2", "val2")

	if want, have := 2, len(logger.kvs); want != have {
		t.Fatalf("unexpected number of key values, want %d, have %d", want, have)
	}
}

func TestWithSuccessWithEmptyContext(t *testing.T) {
	ctx := context.Background()
	logger := &inspectLogger{}

	wLogger := With(ctx, logger)
	wLogger.Log("key2", "val2")

	if want, have := 1, len(logger.kvs); want != have {
		t.Fatalf("unexpected number of key values, want %d, have %d", want, have)
	}
}

func TestClearSuccess(t *testing.T) {
	ctx := Add(context.Background(), "key", "value")
	ctx2 := Clear(ctx)
	kvs := get(ctx2)
	if want, have := 0, len(kvs); want != have {
		t.Fatalf("unexpected number of values, want: %d, have %d", want, have)
	}
}

func TestProvidersSuccess(t *testing.T) {
	RegisterProvider(func(ctx context.Context) context.Context {
		return Add(ctx, "key_0", "value_0")
	})
	ctx := Add(context.Background(), "key_1", "value_1")
	l := &inspectLogger{}
	With(ctx, l).Log("key_2", "value_2")
	if want, have := 3, len(l.kvs); want != have {
		t.Fatalf("unexpected number of values, want: %d, have %d", want, have)
	}
}
