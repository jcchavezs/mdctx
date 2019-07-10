# mdctx Zipkin provider

Providers allows you to include additional context to the logs coming from other
sources.

The zipkin provider allows the user to include the `traceID` and `spanID` as log
fields:

## TraceID only

```go
package main

import "github.com/jcchavezs/mdctx/providers/zipkin"

func main() {
	mdctx.RegisterProvider(zipkin.TraceIDProvider)
}
```

## TraceID and SpanID

```go
package main

import "github.com/jcchavezs/mdctx/providers/zipkin"

func main() {
	mdctx.RegisterProvider(zipkin.TraceNSpanIDProvider)
}
```