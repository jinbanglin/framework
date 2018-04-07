package grpc

import (
	"github.com/jinbanglin/moss"
	"github.com/jinbanglin/moss/log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Handler which should be called from the gRPC binding of the service
// implementation. The incoming request parameter, and returned response
// parameter, are both gRPC types, not user-domain.
type Handler interface {
	ServeGRPC(ctx context.Context, request interface{}) (context.Context, interface{}, error)
}

// Server wraps an endpoint and implements grpc.Handler.
type Server struct {
	e         moss.Endpoint
	dec       DecodeRequestFunc
	enc       EncodeResponseFunc
	before    []ServerRequestFunc
	after     []ServerResponseFunc
	finalizer []ServerFinalizerFunc
	logger    log.Logger
}

// NewServer constructs a new server, which implements wraps the provided
// endpoint and implements the Handler interface. Consumers should write
// bindings that adapt the concrete gRPC methods from their compiled protobuf
// definitions to individual handlers. Request and response objects are from the
// caller business domain, not gRPC request and reply types.
func NewServer(
	e moss.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
) *Server {
	s := &Server{
		e:   e,
		dec: dec,
		enc: enc,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// ServerOption sets an optional parameter for servers.
type ServerOption func(*Server)

// ServerBefore functions are executed on the HTTP request object before the
// request is decoded.
func ServerBefore(before ...ServerRequestFunc) ServerOption {
	return func(s *Server) { s.before = append(s.before, before...) }
}

// ServerAfter functions are executed on the HTTP response writer after the
// endpoint is invoked, but before anything is written to the client.
func ServerAfter(after ...ServerResponseFunc) ServerOption {
	return func(s *Server) { s.after = append(s.after, after...) }
}

// ServerErrorLogger is used to log non-terminal errors. By default, no errors
// are logged.
func ServerErrorLogger(logger log.Logger) ServerOption {
	return func(s *Server) { s.logger = logger }
}

// ServerFinalizer is executed at the end of every gRPC request.
// By default, no finalizer is registered.
func ServerFinalizer(f ...ServerFinalizerFunc) ServerOption {
	return func(s *Server) { s.finalizer = append(s.finalizer, f...) }
}

func (s Server) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	var err error
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	if len(s.finalizer) > 0 {
		defer func() {
			for _, f := range s.finalizer {
				f(ctx, err)
			}
		}()
	}
	for _, f := range s.before {
		ctx = f(ctx, md)
	}
	var request, response, grpcResp interface{}
	request, err = s.dec(ctx, req)
	if err != nil {
		log.Error(err)
		return ctx, response, err
	}
	response, err = s.e(ctx, request)
	if err != nil {
		log.Error(err)
		return ctx, response, err
	}
	var mdHeader, mdTrailer metadata.MD
	for _, f := range s.after {
		ctx = f(ctx, &mdHeader, &mdTrailer)
	}
	grpcResp, err = s.enc(ctx, response)
	if err != nil {
		log.Error(err)
		return ctx, response, err
	}
	if len(mdHeader) > 0 {
		if err = grpc.SendHeader(ctx, mdHeader); err != nil {
			log.Error(err)
			return ctx, response, err
		}
	}
	if len(mdTrailer) > 0 {
		if err = grpc.SetTrailer(ctx, mdTrailer); err != nil {
			log.Error(err)
			return ctx, response, err
		}
	}
	return ctx, grpcResp, nil
}

type ServerFinalizerFunc func(ctx context.Context, err error)

func Interceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	ctx = context.WithValue(ctx, ContextKeyRequestMethod, info.FullMethod)
	return handler(ctx, req)
}
