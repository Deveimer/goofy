package handler

import (
	"encoding/json"
	"github.com/gookit/color"
	"go.opencensus.io/trace"
	"gofy/pkg/types"
	"log"
	"net/http"
	"os"
)

type Handler func(r *http.Request) (interface{}, error)

func (h Handler) HandleHTTP(w http.ResponseWriter, r *http.Request) {
	traceId := trace.FromContext(r.Context()).SpanContext().TraceID.String()
	data, err := h(r)
	var res Response
	switch t := err.(type) {

	case nil:
		res = Response{http.StatusOK, "SUCCESS", data}

	case types.MissingParam, types.InvalidParam:
		res = Response{http.StatusBadRequest, "ERROR", ErrorData{traceId, t, t.Error()}}

	case types.EntityNotFound:
		res = Response{http.StatusNotFound, "ERROR", ErrorData{traceId, t, t.Error()}}

	case types.Response:
		res = Response{Code: t.StatusCode, Status: "ERROR", Data: ErrorData{Id: traceId, Details: t, Message: t.Error()}}

	default:
		// This is unexpected error. So this will always be 500.
		logger := log.New(os.Stderr, color.Red.Render("[ERR] "), 0)
		line, _ := json.Marshal(ErrorData{traceId, t, t.Error()})
		logger.Println(string(line))

		res = Response{http.StatusInternalServerError, "ERROR", ErrorData{traceId, nil, "Internal Server Error"}}
	}

	w.Header().Set("hes-x-trace", traceId)
	w.WriteHeader(res.Code)
	_ = json.NewEncoder(w).Encode(res)

}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorData struct {
	Id      string      `json:"id"`
	Details interface{} `json:"details,omitempty"`
	Message string      `json:"message"`
}
