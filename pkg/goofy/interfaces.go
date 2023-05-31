package goofy

import (
	"net/http"
)

type Router interface {
	http.Handler

	Route(method string, path string, handler Handler)
	Use(middleware ...Middleware)
}
