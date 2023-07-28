package goofy

import (
	"encoding/json"
	"github.com/Deveimer/goofy/pkg/goofy/errors"
	"net/http"
)

type Handler func(ctx *Context) (interface{}, error)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Context().Value(1).(*Context)
	data, err := h(c)
	var res Response
	switch t := err.(type) {

	case nil:
		status := http.StatusOK
		if r.Method == "DELETE" {
			status = http.StatusNoContent
		}
		res = Response{Code: status, Status: "SUCCESS", Data: data}

	case errors.MissingParam, errors.InvalidParam, errors.EntityAlreadyExists:
		res = Response{Code: http.StatusBadRequest, Status: "ERROR", Error: ErrorData{Message: t.Error()}}

	case errors.EntityNotFound:
		res = Response{Code: http.StatusNotFound, Status: "ERROR", Error: ErrorData{Message: t.Error()}}

	case errors.Response:
		res = Response{Code: t.StatusCode, Status: "ERROR", Error: ErrorData{Details: t, Message: t.Error()}}

	default:
		res = Response{Code: http.StatusInternalServerError, Status: "ERROR", Error: ErrorData{nil, "Internal Server Error"}}
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(res.Code)
	_ = json.NewEncoder(w).Encode(res)

}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

type ErrorData struct {
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message"`
}
