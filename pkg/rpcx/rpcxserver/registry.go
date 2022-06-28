package rpcxserver

import (
	"github.com/bufbuild/connect-go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type registryImpl struct {
	handlerMap map[string]http.Handler
	options    []connect.HandlerOption
}

func (s *registryImpl) Options() []connect.HandlerOption {
	return s.options
}

func (s *registryImpl) AddHandler(path string, handler http.Handler) {
	s.handlerMap[path] = handler
}

func (s *registryImpl) registerHttp(server *echo.Echo) {
	for path, handler := range s.handlerMap {
		server.POST(path+"*", echo.WrapHandler(handler))
	}
}
