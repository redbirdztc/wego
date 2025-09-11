package jaeger

import (
	"context"
	"log/slog"
	"net/http"

	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

// Initialize OpenTelemetry TracerProvider and connect to Jaeger
func InitTracerProvider(ctx context.Context, jaegerEndpoint string) error {
	if os.Getenv("TRACE_ON") != "true" {
		slog.Error("OpenTelemetry tracing not enabled, skipping initialization")
		return nil
	}

	// Create Jaeger exporter (sends trace data to Jaeger)
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		slog.Error("Failed to create Jaeger exporter", "err", err)
		return err
	}
	serviceName := os.Getenv("SERVICE_NAME")
	// Configure resource (service name and other metadata)
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName), // Service name (displayed in Jaeger UI)
		),
	)
	if err != nil {
		slog.Error("Failed to create resource", "err", err)
	}

	// Create TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),                // Batch export of trace data (improves performance)
		sdktrace.WithResource(res),                    // Associate resource information
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // Full sampling for development environment
		// For production environment, use probability sampling:
		// sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(0.1))),
	)

	// Set global tracer
	otel.SetTracerProvider(tp)

	// Set context propagator (for passing trace information across services)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, // Standard W3C Trace Context format
		propagation.Baggage{},
	))
	slog.Info("OpenTelemetry TracerProvider initialized successfully", "jaegerEndpoint", jaegerEndpoint, "serviceName", serviceName)
	return nil
}

func TraceContext(req *http.Request) (context.Context, func(options ...trace.SpanEndOption)) {
	if os.Getenv("TRACE_ON") == "true" {
		ctx := otel.GetTextMapPropagator().Extract(req.Context(), propagation.HeaderCarrier(req.Header))
		var span trace.Span
		ctx, span = otel.GetTracerProvider().Tracer("cube-permission").Start(ctx, req.Method+" "+req.URL.Path)
		return ctx, span.End
	}
	return req.Context(), func(options ...trace.SpanEndOption) {}
}
