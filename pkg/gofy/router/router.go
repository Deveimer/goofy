package router

import (
	"github.com/gorilla/mux"
	"github.com/varun-singhh/gofy/pkg/gofy/handler"
	"net/http"
)

type router struct {
	mux.Router

	prefix string
}

type Router interface {
	http.Handler

	Route(method string, path string, handler handler.Handler)
}

func NewRouter() Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	r := router{Router: *muxRouter}

	return &r
}

func (r *router) Route(method, path string, handler handler.Handler) {
	if r.prefix != "" {
		path = r.prefix + path
	}

	r.Router.NewRoute().Methods(method).Path(path).Handler(handler)
}
