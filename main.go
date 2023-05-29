package main

import (
	"github.com/varun-singhh/gofy/pkg/gofy"
	"net/http"
)

func main() {

	app := gofy.New()

	// POST endpoints
	app.Server.HTTP.Port = 2222

	app.GET("/hi", func(r *http.Request) (interface{}, error) {
		return "hello", nil
	})

	app.Start()

}
