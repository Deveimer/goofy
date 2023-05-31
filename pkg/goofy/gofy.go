package goofy

import (
	"net/http"
	"strings"

	"github.com/varun-singhh/gofy/pkg/goofy/config"
	"github.com/varun-singhh/gofy/pkg/goofy/log"
)

type Goofy struct {
	Config config.Config
	Logger log.Logger
	Server *Server
}

func (k *Goofy) Start() {
	k.Server.Start(k.Logger)
}

func (k *Goofy) addRoute(method, path string, handler Handler) {

	if path != "/" {
		path = strings.TrimSuffix(path, "/")
		k.Server.Router.Route(method, path+"/", handler)
	}
	k.Server.Router.Route(method, path, handler)

}

func (k *Goofy) GET(path string, handler Handler) {
	k.addRoute(http.MethodGet, path, handler)
}

func (k *Goofy) PUT(path string, handler Handler) {
	k.addRoute(http.MethodPut, path, handler)
}

func (k *Goofy) POST(path string, handler Handler) {
	k.addRoute(http.MethodPost, path, handler)
}

func (k *Goofy) DELETE(path string, handler Handler) {
	k.addRoute(http.MethodDelete, path, handler)
}

func (k *Goofy) PATCH(path string, handler Handler) {
	k.addRoute(http.MethodPatch, path, handler)
}
