package main

import (
	"fmt"
	"github.com/varun-singhh/gofy/pkg/goofy"
	"github.com/varun-singhh/gofy/pkg/goofy/errors"
)

func main() {

	app := goofy.New()

	// POST endpoints

	app.GET("/hi", func(ctx *goofy.Context) (interface{}, error) {
		fmt.Println(ctx.Request().URL)
		return errors.Response{
			StatusCode: 500,
			Status:     "Internal server error",
			Code:       "500",
		}, nil
	})

	app.Start()

}
