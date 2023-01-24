package middleware

import (
	"blog-service/global"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claims, err := app.ParseToke(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
			if claims.ExpiresAt-time.Now().Unix() < int64(claims.BufferTime) {
				claims.ExpiresAt = time.Now().Add(global.JWTSetting.Expire).Unix()
				tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				newToken, _ := tokenClaims.SignedString(app.GetJWTSecret())
				c.Header("new-token", newToken)
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
