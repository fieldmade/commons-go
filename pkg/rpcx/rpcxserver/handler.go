package rpcxserver

import (
	"github.com/bufbuild/connect-go"
	"github.com/fieldmade/commons-go/pkg/rpcx"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxcodec"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxerr"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxvalidate"
	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

type Handler struct {
	Services             []rpcx.Service
	registry             *registryImpl
	Interceptors         []connect.Interceptor
	JsonMarshalOptions   *protojson.MarshalOptions
	JsonUnmarshalOptions *protojson.UnmarshalOptions
	options              []connect.HandlerOption
}

func (s *Handler) initOptions() {
	if s.JsonMarshalOptions == nil {
		s.JsonMarshalOptions = &protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		}
	}

	if s.JsonUnmarshalOptions == nil {
		s.JsonUnmarshalOptions = &protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}
	}

	s.options = append(s.options, connect.WithCodec(rpcxcodec.NewJsonCodec(s.JsonMarshalOptions, s.JsonUnmarshalOptions)))
	s.options = append(s.options, connect.WithInterceptors(s.allInterceptors()...))
}

func (s *Handler) allInterceptors() []connect.Interceptor {
	var res []connect.Interceptor

	res = append(res, rpcxerr.NewServerInterceptor())
	res = append(res, rpcxvalidate.NewServerInterceptor())

	res = append(res, s.Interceptors...)
	return res
}

func (s *Handler) initRegistry() {
	s.registry = &registryImpl{
		handlerMap: map[string]http.Handler{},
		options:    s.options,
	}

	for _, service := range s.Services {
		service.RegisterRpc(s.registry)
	}
}

func (s *Handler) RegisterHttp(server *echo.Echo) error {
	s.initOptions()
	s.initRegistry()

	s.registry.registerHttp(server)

	return nil
}
