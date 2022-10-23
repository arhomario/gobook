package routes

import (
	apiControllerV1 "gobook/controllers/api/v1"
	"gobook/middlewares"

	"github.com/gin-gonic/gin"
)

func Api(route *gin.Engine) {

	route.POST("api/login", apiControllerV1.Postlogin)

	api := route.Group("/api/v1")
	api.Use(middlewares.SetMiddlewareAuthentication())
	{
		api.GET("books", apiControllerV1.GetBooks)
		api.POST("books", apiControllerV1.StoreBook)
		api.GET("books/:id", apiControllerV1.EditBook)
		api.PATCH("books/:id", apiControllerV1.UpdateBook)
		api.DELETE("books/:id", apiControllerV1.DeleteBook)
	}

}
