package main

import (
	"gobook/config"
	"gobook/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db := config.Database()
	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	routes.Api(r)
	routes.Web(r)
	r.Run(":8080")
}
