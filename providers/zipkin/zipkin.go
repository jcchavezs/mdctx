package zipkinmdctx

import (
	"context"

	"github.com/openzipkin/zipkin-go"
)

var KVProvider = func(ctx context.Context) []interface{} {
	span := zipkin.SpanFromContext(ctx)
	if span != nil {
		if traceID := span.Context().ID.String(); traceID != "" {
			return []interface{}{"trace_id", traceID}
		}
	}
	return nil
}
