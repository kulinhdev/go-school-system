package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kulinhdev/serentyspringsedu/api/res"
	controllers "github.com/kulinhdev/serentyspringsedu/controller"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/middlewares"
	"github.com/kulinhdev/serentyspringsedu/models"
	"net/http"
)

func Initialize() {
	router := gin.Default()
	// Config Cors
	corConfig := cors.DefaultConfig()
	corConfig.AllowOrigins = []string{"*"}
	corConfig.AllowHeaders = []string{
		"authorization", "Authorization",
		"content-type", "accept",
		"referer", "user-agent",
	}
	router.Use(cors.New(corConfig))

	// Get all routes
	studentController := controllers.NewStudentController(models.DB)
	authController := controllers.NewAuthController(models.DB)

	// Register all routes
	apiRoute := router.Group("/api")
	{

		studentRoute := apiRoute.Group("/students")
		studentRoute.Use(middlewares.DeserializeUser())
		{
			studentRoute.GET("/", studentController.List)
			studentRoute.POST("/", studentController.Create)
			studentRoute.GET("/:id", studentController.FindById)
			studentRoute.PUT("/:id", studentController.Update)
			studentRoute.DELETE("/:id", studentController.Delete)
		}

		authRoute := apiRoute.Group("auth")
		{
			authRoute.POST("/register", authController.RegisterUser)
			authRoute.POST("/login", authController.LoginUser)
			authRoute.GET("/refresh", middlewares.DeserializeUser(), authController.RefreshAccessToken)
			authRoute.POST("/logout", middlewares.DeserializeUser(), authController.LogoutUser)
		}
	}

	router.NoRoute(func(ctx *gin.Context) {
		res.ResponseError(ctx, http.StatusNotFound, "Page not found", nil)
	})

	// Listen port:9000
	err := router.Run(fmt.Sprintf(":%s", initializers.Config.RunPort))
	if err != nil {
		fmt.Printf("Initialize routes failed: %v", err)
	}
}
