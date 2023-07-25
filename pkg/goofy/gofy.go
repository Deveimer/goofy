package goofy

import (
	"net/http"
	"strings"

	"gorm.io/gorm"

	"github.com/Deveimer/goofy/pkg/goofy/config"
	"github.com/Deveimer/goofy/pkg/goofy/log"
)

type Goofy struct {
	Config   config.Config
	Logger   log.Logger
	Server   *Server
	Database *gorm.DB
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
