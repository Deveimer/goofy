package goofy

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type router struct {
	mux.Router
	prefix string
}

func NewRouter() Router {
	muxRouter := mux.NewRouter().StrictSlash(false)
	r := router{Router: *muxRouter}
	return &r
}

func (r *router) Route(method, path string, handler Handler) {
	if r.prefix != "" {
		path = r.prefix + path
	}

	r.Router.NewRoute().Methods(method).Path(path).Handler(handler)
}

func (r *router) Use(middleware ...Middleware) {
	mwf := make([]mux.MiddlewareFunc, 0, len(middleware))
	for _, m := range middleware {
		mwf = append(mwf, mux.MiddlewareFunc(m))
	}

	r.Router.Use(mwf...)
}
