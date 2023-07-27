package goofy

import (
	"context"
	"github.com/Deveimer/goofy/pkg/goofy/log"
	"github.com/Deveimer/goofy/pkg/goofy/request"
	"github.com/Deveimer/goofy/pkg/goofy/response"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func NewContext(w response.Responder, r request.Request, k *Goofy) *Context {
	return &Context{
		req:    r,
		res:    w,
		Goofy:  k,
		Logger: log.NewLogger(),
	}
}

type Context struct {
	context.Context
	*Goofy
	Logger log.Logger
	res    response.Responder
	req    request.Request
}

func (c *Context) Reset(w response.Responder, r request.Request) {
	c.req = r
	c.res = w
	c.Context = nil
	c.Logger = log.NewLogger()
}

func (c *Context) Request() *http.Request {
	return c.req.Request()
}

func (c *Context) Param(key string) string {
	return c.req.Param(key)
}

func (c *Context) Params() map[string]string {
	return c.req.Params()
}

func (c *Context) PathParam(key string) string {
	return c.req.PathParam(key)
}

func (c *Context) Bind(i interface{}) error {
	return c.req.Bind(i)
}

func (c *Context) BindStrict(i interface{}) error {
	return c.req.BindStrict(i)
}

func (c *Context) Header(key string) string {
	return c.req.Header(key)
}

// Log logs the key-value pair into the logs
func (c *Context) Log(key string, value interface{}) {
	// This section takes care of middleware logging
	if key == "correlationID" { // This condition will not allow the user to unset the CorrelationID.
		return
	}

	r := c.Request()
	appLogData, ok := r.Context().Value("appLogData").(*sync.Map)

	if !ok {
		c.Logger.Warn("couldn't log appData")
		return
	}

	appLogData.Store(key, value)
	*r = *r.Clone(context.WithValue(r.Context(), "appLogData", appLogData))

	// This section takes care of all the individual context loggers
	c.Logger.AddData(key, value)
}

// SetPathParams sets the URL path variables to the given value. These can be accessed
// by c.PathParam(key). This method should only be used for testing purposes.
func (c *Context) SetPathParams(pathParams map[string]string) {
	r := c.req.Request()

	r = mux.SetURLVars(r, pathParams)

	c.req = request.NewHTTPRequest(r)
}
