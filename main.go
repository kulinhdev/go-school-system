package main

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/kulinhdev/serentyspringsedu/api"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/models"
)

// Custom validator
func evenAge(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	return age > 0 && age%2 == 0
}

func main() {
	// Load Config
	initializers.Initialize()

	// Init DB
	models.Initialize()

	// Register custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("evenAge", evenAge)
	}

	// Setup API routes
	api.Initialize()

}
