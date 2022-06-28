package rpcx

import (
	"github.com/bufbuild/connect-go"
	"net/http"
)

type Service interface {
	RegisterRpc(reg Registry)
}

type Registry interface {
	Options() []connect.HandlerOption
	AddHandler(path string, handler http.Handler)
}
