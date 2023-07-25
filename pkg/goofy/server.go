package goofy

import (
	"context"
	"github.com/Deveimer/goofy/pkg/goofy/log"
	"github.com/Deveimer/goofy/pkg/goofy/request"
	"github.com/Deveimer/goofy/pkg/goofy/response"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	Router      Router
	done        chan bool
	HTTP        HTTP
	contextPool sync.Pool
}

type HTTP struct {
	Port int
}

func NewServer(goofy *Goofy) *Server {
	s := &Server{
		Router: NewRouter(),
		HTTP:   HTTP{},
	}

	s.contextPool.New = func() interface{} {
		return NewContext(nil, nil, goofy)
	}

	return s
}

func (s *Server) Start(logger log.Logger) {

	var srv *http.Server

	s.Router.Use(s.contextInjector)

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

func (s *Server) contextInjector(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := s.contextPool.Get().(*Context)
		c.Reset(response.NewContextualResponder(w, r), request.NewHTTPRequest(r))
		*r = *r.Clone(context.WithValue(r.Context(), "appLogData", &sync.Map{}))
		c.Context = r.Context()
		*r = *r.Clone(context.WithValue(c.Context, 1, c))

		inner.ServeHTTP(w, r)

		s.contextPool.Put(c)
	})
}
