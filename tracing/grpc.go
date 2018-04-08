package tracing

import (
	"golang.org/x/net/context"
	"encoding/base64"
	"strings"

	"github.com/jinbanglin/moss/log"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc/metadata"
)

// ContextToGRPC returns a grpc RequestFunc that injects an OpenTracing Span
// found in `ctx` into the grpc Metadata. If no such Span can be found, the
// RequestFunc is a noop.
func ContextToGRPC(tracer opentracing.Tracer) func(ctx context.Context, md *metadata.MD) context.Context {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		if span := opentracing.SpanFromContext(ctx); span != nil {
			// There's nothing we can do with an error here.
			if err := tracer.Inject(span.Context(), opentracing.TextMap, metadataReaderWriter{md}); err != nil {
				log.Error(err)
			}
		}
		return ctx
	}
}

// GRPCToContext returns a grpc RequestFunc that tries to join with an
// OpenTracing trace found in `req` and starts a new Span called
// `operationName` accordingly. If no trace could be found in `req`, the Span
// will be a trace root. The Span is incorporated in the returned Context and
// can be retrieved with opentracing.SpanFromContext(ctx).
func GRPCToContext(tracer opentracing.Tracer, operationName string) func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		var span opentracing.Span
		wireContext, err := tracer.Extract(opentracing.TextMap, metadataReaderWriter{&md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			log.Error(err)
		}
		span = tracer.StartSpan(operationName, ext.RPCServerOption(wireContext))
		return opentracing.ContextWithSpan(ctx, span)
	}
}

// A type that conforms to opentracing.TextMapReader and
// opentracing.TextMapWriter.
type metadataReaderWriter struct {
	*metadata.MD
}

func (w metadataReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	if strings.HasSuffix(key, "-bin") {
		val = string(base64.StdEncoding.EncodeToString([]byte(val)))
	}
	(*w.MD)[key] = append((*w.MD)[key], val)
}

func (w metadataReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range *w.MD {
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
