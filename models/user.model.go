package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	StudentCode string  `gorm:"type:varchar(100);unique_index"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	Username    string  `gorm:"size:255;not null;unique" json:"username"`
	Password    string  `gorm:"size:255;not null;" json:"-"`
	Phone       *string `gorm:"type:varchar(100);unique;index"`
	Email       string  `gorm:"type:varchar(100);unique"`
	Photo       string
	Status      int // 0: Inactive, 1: Active
	Gender      int // 0: Female, 1: Male
	Role        int // 0: Admin, 1: User
	Birthday    time.Time
	Address     *string `gorm:"type:text"`
}

type UserResponse struct {
	ID        uint
	Name      string
	Username  string
	Phone     *string
	Email     string
	Photo     string
	Gender    int
	Birthday  time.Time
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegister struct {
	Name            string
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Photo           string
}

type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
