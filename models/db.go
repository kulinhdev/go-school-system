package models

import (
	"fmt"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Initialize() {
	// Initialize the database connection
	dsn := initializers.GetDatabaseURL()
	ConnectDB(dsn)

	// Migrate DB
	DBAutoMigration()
}

func ConnectDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database!")
	}
	fmt.Println("🚀 Connected Successfully to the Database!")
}

func DBAutoMigration() {
	fmt.Sprintln("Starting migration ...")
	err := DB.AutoMigrate(&Student{}, &User{})
	if err != nil {
		fmt.Printf("Migration Failed: %v\n", err)
		return
	}
	fmt.Sprintln("Migration Successfully!")
}
