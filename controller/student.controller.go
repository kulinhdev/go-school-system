package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kulinhdev/serentyspringsedu/api/res"
	"github.com/kulinhdev/serentyspringsedu/helpers"
	"github.com/kulinhdev/serentyspringsedu/models"
	"gorm.io/gorm"
	"net/http"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(DB *gorm.DB) StudentController {
	return StudentController{DB}
}

// List retrieves a list of students.
// @Summary Retrieve a list of students
// @ID list-students
// @Tags Students
// @Accept json
// @Produce json
// @Success 200 {array} models.Student
// @Router /api/students [get]
func (ctl *StudentController) List(ctx *gin.Context) {
	var students []models.Student
	if err := ctl.DB.Find(&students).Error; err != nil {
		res.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	res.ResponseSuccess(ctx, http.StatusOK, students)
}

// Create creates a new student.
// @Summary Create a new student
// @ID create-student
// @Tags Students
// @Accept json
// @Produce json
// @Param student body models.Student true "Student data"
// @Success 201 {object} models.Student
// @Router /api/students [post]
func (ctl *StudentController) Create(ctx *gin.Context) {
	var student models.Student
	var ve validator.ValidationErrors
	if err := ctx.ShouldBindBodyWith(&student, binding.JSON); err != nil {
		ve := helpers.CustomMessageErrors(err, ve)
		res.ResponseError(ctx, http.StatusBadRequest, err.Error(), gin.H{"errors": ve})
		return
	}

	if err := ctl.DB.Create(&student).Error; err != nil {
		res.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	res.ResponseSuccess(ctx, http.StatusCreated, student)
}

// FindById retrieves a student by ID.
// @Summary Retrieve a student by ID
// @ID find-student-by-id
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Router /api/students/{id} [get]
func (ctl *StudentController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student
	if err := ctl.DB.First(&student, "id = ?", id).Error; err != nil {
		res.ResponseError(ctx, http.StatusNotFound, "Record not found", nil)
		return
	}
	res.ResponseSuccess(ctx, http.StatusOK, student)
}

// Update updates a student by ID.
// @Summary Update a student by ID
// @ID update-student-by-id
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Student data"
// @Success 200 {object} models.Student
// @Router /api/students/{id} [put]
func (ctl *StudentController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student
	if err := ctl.DB.First(&student, "id = ?", id).Error; err != nil {
		res.ResponseError(ctx, http.StatusNotFound, "Record not found", nil)
		return
	}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		res.ResponseError(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctl.DB.Save(&student).Error; err != nil {
		res.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	res.ResponseSuccess(ctx, http.StatusOK, student)
}

// Delete deletes a student by ID.
// @Summary Delete a student by ID
// @ID delete-student-by-id
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Success 204
// @Router /api/students/{id} [delete]
func (ctl *StudentController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student
	if err := ctl.DB.First(&student, "id = ?", id).Error; err != nil {
		res.ResponseError(ctx, http.StatusNotFound, "Record not found", nil)
		return
	}

	if err := ctl.DB.Delete(&student).Error; err != nil {
		res.ResponseError(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	res.ResponseSuccess(ctx, http.StatusNoContent, nil)
}
