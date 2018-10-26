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
	if want, have := "key", kvs[0]; want != have {
		t.Fatalf("unexpected key, want: %q, have: %q", want, have)
	}

	if want, have := "val", kvs[1]; want != have {
		t.Fatalf("unexpected value, want: %q, have: %q", want, have)
	}
}

type inspectLogger struct {
	kvs keyVals
}

func (l *inspectLogger) Log(kvs ...interface{}) error {
	kvs = append(l.kvs, kvs...)
	l.kvs = keyVals(kvs)
	return nil
}

func TestWithSuccessWhenHavingValues(t *testing.T) {
	ctx := Add(context.Background(), "key", "val")
	logger := &inspectLogger{}

	wLogger := With(ctx, logger)
	wLogger.Log("key2", "val2")

	if want, have := 4, len(logger.kvs); want != have {
		t.Fatalf("unexpected number of key values, want %d, have %d", want, have)
	}
}

func TestWithSuccessWithEmptyContext(t *testing.T) {
	ctx := context.Background()
	logger := &inspectLogger{}

	wLogger := With(ctx, logger)
	wLogger.Log("key2", "val2")

	if want, have := 2, len(logger.kvs); want != have {
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
