package zipkinmdctx

import (
	"context"

	"github.com/jcchavezs/mdctx"

	"github.com/openzipkin/zipkin-go"
)

// Provider allows users to inject zipkin's traceID value into
// the log record values.
var Provider = func(ctx context.Context) context.Context {
	span := zipkin.SpanFromContext(ctx)
	if span != nil {
		if traceID := span.Context().ID.String(); traceID != "" {
			return mdctx.Add(ctx, "trace_id", traceID)
		}
	}
	return ctx
}
