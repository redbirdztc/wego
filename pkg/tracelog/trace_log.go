package tracelog

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

// TraceHandler wraps another slog.Handler and adds trace ID to log records
type TraceHandler struct {
	handler slog.Handler
}

// NewTraceHandler creates a new TraceHandler that wraps the given handler
func NewTraceHandler(handler slog.Handler) *TraceHandler {
	return &TraceHandler{handler: handler}
}

// Enabled implements slog.Handler
func (h *TraceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle implements slog.Handler and adds trace ID to the record
func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	// Extract trace ID from context
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		traceID := span.SpanContext().TraceID().String()
		spanID := span.SpanContext().SpanID().String()

		// Add trace ID and span ID as attributes
		r.Add("trace_id", traceID)
		r.Add("span_id", spanID)
	}

	return h.handler.Handle(ctx, r)
}

// WithAttrs implements slog.Handler
func (h *TraceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TraceHandler{handler: h.handler.WithAttrs(attrs)}
}

// WithGroup implements slog.Handler
func (h *TraceHandler) WithGroup(name string) slog.Handler {
	return &TraceHandler{handler: h.handler.WithGroup(name)}
}

func InitSlog() {
	// Create JSON handler with pretty options
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Rename the default timestamp key to be more readable
			if a.Key == slog.TimeKey {
				a.Key = "timestamp"
			}
			return a
		},
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, opts)

	// Wrap with trace handler to automatically include trace ID
	traceHandler := NewTraceHandler(jsonHandler)

	jsonSlog := slog.New(traceHandler)

	slog.SetDefault(jsonSlog)
}
