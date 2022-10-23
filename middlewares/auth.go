package middlewares

import (
	"gobook/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetMiddlewareAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := auth.ExtractToken(ctx)
		if token == "" {
			code = 4001
		} else {
			err := auth.TokenValid(token)
			if err != nil {
				ctx.Abort()
				code = 4004
			}
		}

		if code != 200 {
			ctx.JSON(http.StatusOK, gin.H{
				"Code":    4004,
				"Message": "Token Invalid",
				"data":    data,
			})

			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
