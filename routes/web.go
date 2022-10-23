package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Web(route *gin.Engine) {
	api := route.Group("/web")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": " Web",
			})
		})
	}
}
