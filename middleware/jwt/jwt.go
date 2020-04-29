package jwt 

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"

	"Pet/pkg/util"
	"Pet/constants"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constants.SUCCESS
		token := c.Query("token")
		if token = "" {
			code := constants.INVALID_PARAMS
		} else {
			claims,err := util.ParseToken(token)
			if err != nil {
				code = constatns.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = constants.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != constants.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": constants.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}