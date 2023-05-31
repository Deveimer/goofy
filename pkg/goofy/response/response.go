package response

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/varun-singhh/gofy/pkg/goofy/errors"
	"net/http"
	"strings"
)

type responseType int

const (
	JSON responseType = iota
	XML
	TEXT
)

type HTTP struct {
	path    string
	method  string
	w       http.ResponseWriter
	resType responseType
}

func NewContextualResponder(w http.ResponseWriter, r *http.Request) Responder {
	route := mux.CurrentRoute(r)

	var path string
	if route != nil {
		path, _ = route.GetPathTemplate()
		// remove the trailing slash
		path = strings.TrimSuffix(path, "/")
	}

	responder := &HTTP{
		w:      w,
		method: r.Method,
		path:   path,
	}

	cType := r.Header.Get("Content-type")
	switch cType {
	case "text/xml", "application/xml":
		responder.resType = XML
	case "text/plain":
		responder.resType = TEXT
	default:
		responder.resType = JSON
	}

	return responder
}

func (h HTTP) Response(data interface{}, err error) {
	var (
		response   interface{}
		statusCode int
	)

	response = data
	statusCode = getStatusCode(h.method, data, err)

	h.processResponse(statusCode, response)
}

func (h HTTP) processResponse(statusCode int, response interface{}) {
	switch h.resType {
	case JSON:
		h.w.Header().Set("Content-type", "application/json")
		h.w.WriteHeader(statusCode)

		if response != nil {
			_ = json.NewEncoder(h.w).Encode(response)
		}

	case XML:
		h.w.Header().Set("Content-type", "application/xml")
		h.w.WriteHeader(statusCode)

		if response != nil {
			_ = xml.NewEncoder(h.w).Encode(response)
		}
	case TEXT:
		h.w.Header().Set("Content-type", "text/plain")
		h.w.WriteHeader(statusCode)

		if response != nil {
			_, _ = fmt.Fprintf(h.w, "%s", response)
		}
	}
}

func getStatusCode(method string, data interface{}, err error) int {
	statusCode := 200

	if err == nil {
		if method == http.MethodPost {
			statusCode = 201
		} else if method == http.MethodDelete {
			statusCode = 204
		}

		return statusCode
	}

	switch t := err.(type) {
	case errors.MissingParam, errors.InvalidParam:
		statusCode = 400
	case errors.EntityNotFound:
		statusCode = 404
	case errors.Response:
		statusCode = t.StatusCode
	default:
		statusCode = 500
	}

	return statusCode
}
