package request

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
)

type HTTP struct {
	req        *http.Request
	pathParams map[string]string
}

// NewHTTPRequest injects a *http.Request into a gofr.Request variable
func NewHTTPRequest(r *http.Request) Request {
	return &HTTP{
		req: r,
	}
}

// Request returns the underlying HTTP request
func (h *HTTP) Request() *http.Request {
	return h.req
}

// String satisfies the Stringer interface for the HTTP type
func (h *HTTP) String() string {
	return fmt.Sprintf("%s %s", h.req.Method, h.req.URL)
}

// Method returns the HTTP Method of the current request
func (h *HTTP) Method() string {
	return h.req.Method
}

// URI returns the current HTTP requests URL-PATH
func (h *HTTP) URI() string {
	return h.req.URL.Path
}

// Param returns the query parameter value for the given key, if any
func (h *HTTP) Param(key string) string {
	values := h.req.URL.Query()[key]
	return strings.Join(values, ",")
}

// ParamNames returns the list of query parameters (keys) for the current request
func (h *HTTP) ParamNames() []string {
	var ( //nolint:prealloc // Preallocating fails a testcase.
		names []string
		q     = h.req.URL.Query()
	)

	for name := range q {
		names = append(names, name)
	}

	return names
}

// Params returns the query parameters for the current request in the form of a mapping of key to it's values as
// comma separated values
func (h *HTTP) Params() map[string]string {
	res := make(map[string]string)
	for key, values := range h.req.URL.Query() {
		res[key] = strings.Join(values, ",")
	}

	return res
}

// PathParam returns the route values for the given key, if any
func (h *HTTP) PathParam(key string) string {
	if h.pathParams == nil {
		h.pathParams = make(map[string]string)
		h.pathParams = mux.Vars(h.req)
	}

	return h.pathParams[key]
}

// Body returns the request body, throws an error if there was one
func (h *HTTP) Body() ([]byte, error) {
	bodyBytes, err := io.ReadAll(h.req.Body)
	if err != nil {
		return nil, err
	}

	h.req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

// Header returns the value associated with the `key`, from the request headers
func (h *HTTP) Header(key string) string {
	return h.req.Header.Get(key)
}

// Bind checks the Content-Type to select a binding encoding automatically.
// Depending on the "Content-Type" header different bindings are used:
// XML binding is used in cae of: "application/xml" or "text/xml"
// JSON binding is used by default
// It decodes the json payload into the type specified as a pointer.
func (h *HTTP) Bind(i interface{}) error {
	body, err := h.Body()
	if err != nil {
		return err
	}

	cType := h.req.Header.Get("Content-type")
	switch cType {
	case "text/xml", "application/xml":
		return xml.Unmarshal(body, &i)
	default:
		return json.Unmarshal(body, &i)
	}
}

func (h *HTTP) BindStrict(i interface{}) error {
	body, err := h.Body()
	if err != nil {
		return err
	}

	cType := h.req.Header.Get("Content-type")
	switch cType {
	case "text/xml", "application/xml":
		return xml.Unmarshal(body, &i)
	default:
		dec := json.NewDecoder(h.req.Body)
		dec.DisallowUnknownFields()

		return dec.Decode(&i)
	}
}
