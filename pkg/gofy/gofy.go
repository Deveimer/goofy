package gofy

import (
	"net/http"
	"strings"

	"github.com/varun-singhh/gofy/pkg/gofy/config"
	"github.com/varun-singhh/gofy/pkg/gofy/handler"
	"github.com/varun-singhh/gofy/pkg/gofy/log"
	"github.com/varun-singhh/gofy/pkg/gofy/server"
)

type Gofy struct {
	Config config.Config
	Logger log.Logger
	Server *server.Server
}

func (k *Gofy) Start() {
	k.Server.Start(k.Logger)
}

func (k *Gofy) addRoute(method, path string, handler handler.Handler) {

	if path != "/" {
		path = strings.TrimSuffix(path, "/")
		k.Server.Router.Route(method, path+"/", handler)
	}
	k.Server.Router.Route(method, path, handler)

}

func (k *Gofy) GET(path string, handler handler.Handler) {
	k.addRoute(http.MethodGet, path, handler)
}

func (k *Gofy) PUT(path string, handler handler.Handler) {
	k.addRoute(http.MethodPut, path, handler)
}

func (k *Gofy) POST(path string, handler handler.Handler) {
	k.addRoute(http.MethodPost, path, handler)
}

func (k *Gofy) DELETE(path string, handler handler.Handler) {
	k.addRoute(http.MethodDelete, path, handler)
}

func (k *Gofy) PATCH(path string, handler handler.Handler) {
	k.addRoute(http.MethodPatch, path, handler)
}
