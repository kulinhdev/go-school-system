package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName string `form:"firstname" json:"firstname" binding:"required"`
	LastName  string `form:"lastname" json:"lastname" binding:"required"`
	Age       int    `form:"age" json:"age" binding:"required,min=16,evenAge"`
}
