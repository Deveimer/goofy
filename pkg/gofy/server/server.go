package server

import (
	"fmt"
	"github.com/varun-singhh/gofy/pkg/gofy/log"
	"github.com/varun-singhh/gofy/pkg/gofy/router"
	"net/http"
	"strconv"
)

type Server struct {
	Router router.Router
	done   chan bool
	HTTP   HTTP
}

type HTTP struct {
	Port int
}

func NewServer() *Server {
	s := &Server{
		Router: router.NewRouter(),
		HTTP:   HTTP{},
	}

	return s
}

func (s *Server) Start(logger log.Logger) {

	logger.Log(fmt.Sprint(s.Router))

	var srv *http.Server

	go func() {
		var err error

		addr := ":" + strconv.Itoa(s.HTTP.Port)
		logger.Logf("starting http server at %s", addr)
		srv = &http.Server{
			Addr:    addr,
			Handler: s.Router,
		}
		err = srv.ListenAndServe()

		if err != nil {
			s.done <- true
			logger.Errorf("error in starting http server at %v: %s", s.HTTP.Port, err)
		}
	}()

	<-s.done
	logger.Log("Server Stopping")
}
