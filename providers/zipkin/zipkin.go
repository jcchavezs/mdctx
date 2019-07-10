package zipkinmdctx

import (
	"context"

	"github.com/jcchavezs/mdctx"

	"github.com/openzipkin/zipkin-go"
)

// TraceIDProvider allows users to inject zipkin's traceID value into
// the log record values.
func TraceIDProvider(ctx context.Context) context.Context {
	span := zipkin.SpanFromContext(ctx)
	if span == nil {
		return ctx
	}

	return mdctx.Add(ctx, "trace_id", span.Context().TraceID.String())
}

// TraceNSpanIDProvider allows users to inject zipkin's traceID and spanID
// value into the log record values.
func TraceNSpanIDProvider(ctx context.Context) context.Context {
	span := zipkin.SpanFromContext(ctx)
	if span == nil {
		return ctx
	}

	return mdctx.Add(
		ctx,
		"trace_id", span.Context().TraceID.String(),
		"span_id", span.Context().ID.String(),
	)
}
