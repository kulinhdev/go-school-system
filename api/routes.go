package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/kulinhdev/serentyspringsedu/controller"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/models"
)

func registerAPIRoutes(router *gin.Engine) {
	// Get all routes
	studentController := controllers.NewStudentController(models.DB)
	routeStudent := NewRouteStudent(studentController)
	authController := controllers.NewAuthController(models.DB)
	routeAuth := NewRouteAuth(authController)

	// Register routes
	apiRoute := router.Group("/api")
	routeStudent.Routes(apiRoute)
	routeAuth.Routes(apiRoute)
}

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

	// Register all routes
	registerAPIRoutes(router)

	// Listen port:9000
	err := router.Run(fmt.Sprintf(":%s", initializers.Config.RunPort))
	if err != nil {
		fmt.Printf("Initialize routes failed: %v", err)
	}
}
