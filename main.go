package main

import (
	"fmt"
	"github.com/varun-singhh/gofy/pkg/goofy"
)

func main() {

	app := goofy.New()

	// POST endpoints
	app.GET("/ping", func(ctx *goofy.Context) (interface{}, error) {
		fmt.Println(ctx.Request().URL)
		return goofy.Response{
			Code:   200,
			Status: "SUCCESS",
			Data:   "pong",
		}, nil
	})

	app.Start()

}
