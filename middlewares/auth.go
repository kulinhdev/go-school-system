package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kulinhdev/serentyspringsedu/helpers"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/models"
	"net/http"
	"strings"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		sub, err := helpers.ValidateToken(access_token, initializers.Config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		result := models.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
