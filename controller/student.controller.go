package controllers

import (
	"github.com/gin-gonic/gin"
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
	ctl.DB.Find(&students)
	ctx.JSON(http.StatusOK, students)
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
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctl.DB.Create(&student)
	ctx.JSON(http.StatusCreated, student)
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
	if err := ctl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, student)
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
	if err := ctl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctl.DB.Save(&student)
	ctx.JSON(http.StatusOK, student)
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
	if err := ctl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctl.DB.Delete(&student)
	ctx.JSON(http.StatusNoContent, gin.H{"success": "Delete record success!"})
}
