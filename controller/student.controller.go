package controllers

import (
	"fmt"
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

// [GET] /api/students
func (crl *StudentController) List(c *gin.Context) {
	var students []models.Student
	crl.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// [POST] /api/students
func (crl *StudentController) Create(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	crl.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

// [GET] /api/students/:id
func (crl *StudentController) FindById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := crl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Record with Id=%s not found", id)})
		return
	}
	c.JSON(http.StatusOK, student)
}

// [PUT] /api/students/:id
func (crl *StudentController) Update(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := crl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crl.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

// [DELETE] /api/students/:id
func (crl *StudentController) Delete(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := crl.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	crl.DB.Delete(&student)
	c.JSON(http.StatusNoContent, gin.H{"success": "Delete record success!"})
}
