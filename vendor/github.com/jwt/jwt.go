/*
#.____    .__               .__  .__       .__     __             
#|    |   |__| _____   ____ |  | |__| ____ |  |___/  |_           
#|    |   |  |/     \_/ __ \|  | |  |/ ___\|  |  \   __\          
#|    |___|  |  Y Y  \  ___/|  |_|  / /_/  >   Y  \  |            
#|_______ \__|__|_|  /\___  >____/__\___  /|___|  /__|            
#        \/        \/     \/       /_____/      \/                
#        _______          __                       __             
#        \      \   _____/  |___  _  _____________|  | __  ______ 
#        /   |   \_/ __ \   __\ \/ \/ /  _ \_  __ \  |/ / /  ___/ 
#       /    |    \  ___/|  |  \     (  <_> )  | \/    <  \___ \  
#       \____|__  /\___  >__|   \/\_/ \____/|__|  |__|_ \/____  > 
#               \/     \/                              \/     \/  
# ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------
#
# => File: jwt.go
# => Version: X.X.X.X
# => Author: JWT [opensource]
# => Purpose: 
#    This set of middleware services for authentication
*/

package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"command-delivery/pkg/e"
	"command-delivery/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
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

		c.Next()
	}
}
