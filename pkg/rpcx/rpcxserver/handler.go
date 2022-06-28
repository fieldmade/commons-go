package rpcxserver

import (
	"github.com/bufbuild/connect-go"
	"github.com/fieldmade/commons-go/pkg/rpcx"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxcodec"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxerr"
	"github.com/fieldmade/commons-go/pkg/rpcx/rpcxvalidate"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Services     []rpcx.Service
	registry     *registryImpl
	Interceptors []connect.Interceptor
	options      []connect.HandlerOption
}

func (s *Handler) initOptions() {
	s.options = append(s.options, connect.WithCodec(rpcxcodec.NewJsonCodec()))
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
