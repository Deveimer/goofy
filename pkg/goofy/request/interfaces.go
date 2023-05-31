package request

import "net/http"

type Request interface {
	Request() *http.Request
	Params() map[string]string
	Param(string) string
	PathParam(string) string
	Bind(interface{}) error
	BindStrict(interface{}) error
	Header(string) string
}
