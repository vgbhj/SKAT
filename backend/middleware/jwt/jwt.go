package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vgbhj/SKAT/pkg/e"
	"github.com/vgbhj/SKAT/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token, err := c.Cookie("token")
		var claims *util.Claims
		if err != nil {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else if token == "" {
			code = e.INVALID_PARAMS
		} else {
			var err error
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Set("currentUser", claims.Username)
		// c.Set("currentUserId", claims.Username)
		c.Next()
	}
}
