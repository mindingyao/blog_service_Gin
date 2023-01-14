package middleware

import (
	"blog-service/global"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(ctx, "panic recover err: %v", err)
				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
