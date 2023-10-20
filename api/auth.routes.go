package api

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/kulinhdev/serentyspringsedu/controller"
	"github.com/kulinhdev/serentyspringsedu/middlewares"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewRouteAuth(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ctl *AuthRouteController) Routes(rg *gin.RouterGroup) {
	router := rg.Group("auth")
	{
		router.POST("/register", ctl.authController.RegisterUser)
		router.POST("/login", ctl.authController.LoginUser)
		router.GET("/refresh", ctl.authController.RefreshAccessToken)
		router.GET("/logout", middlewares.DeserializeUser(), ctl.authController.LogoutUser)
	}

}
