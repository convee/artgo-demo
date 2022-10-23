package main

import (
	"net/http"

	"github.com/convee/artgo"
)

func main() {
	r := artgo.Default()
	r.GET("/", func(context *artgo.Context) {
		context.String(http.StatusOK, "%s%s", "hello ", " world")
	})

	userG := r.Group("/user")
	userG.Use(func(context *artgo.Context) {
		context.String(http.StatusOK, "route group middleware front\n")
		context.Next()
		context.String(http.StatusOK, "route group middleware after\n")
	})
	{
		userG.GET("/info", func(context *artgo.Context) {
			context.String(http.StatusOK, "user info\n")
		})
	}

	r.Run(":9999")
}
