package goofy

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"
	"github.com/varun-singhh/gofy/pkg/goofy/errors"
)

type Handler func(ctx *Context) (interface{}, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value(1).(*Context)
	data, err := h(c)
	var res Response
	switch t := err.(type) {

	case nil:
		res = Response{http.StatusOK, "SUCCESS", data}

	case errors.MissingParam, errors.InvalidParam:
		res = Response{http.StatusBadRequest, "ERROR", ErrorData{t, t.Error()}}

	case errors.EntityNotFound:
		res = Response{http.StatusNotFound, "ERROR", ErrorData{t, t.Error()}}

	case errors.Response:
		res = Response{Code: t.StatusCode, Status: "ERROR", Data: ErrorData{Details: t, Message: t.Error()}}

	default:
		// This is unexpected error. So this will always be 500.
		logger := log.New(os.Stderr, color.Red.Render("[ERR] "), 0)
		line, _ := json.Marshal(ErrorData{t, t.Error()})
		logger.Println(string(line))

		res = Response{http.StatusInternalServerError, "ERROR", ErrorData{nil, "Internal Server Error"}}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(res.Code)
	_ = json.NewEncoder(w).Encode(res)

}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorData struct {
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message"`
}
