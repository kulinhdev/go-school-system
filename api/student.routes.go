package api

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/kulinhdev/serentyspringsedu/controller"
)

type StudentRouteController struct {
	postController controllers.StudentController
}

func NewRouteStudent(postController controllers.StudentController) StudentRouteController {
	return StudentRouteController{postController}
}

func (ctl *StudentRouteController) Routes(r *gin.RouterGroup) {
	router := r.Group("/students")
	{
		router.GET("/", ctl.postController.List)
		router.POST("/", ctl.postController.Create)
		router.GET("/:id", ctl.postController.FindById)
		router.PUT("/:id", ctl.postController.Update)
		router.DELETE("/:id", ctl.postController.Delete)
	}
}
