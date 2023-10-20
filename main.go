package main

import (
	"github.com/kulinhdev/serentyspringsedu/api"
	"github.com/kulinhdev/serentyspringsedu/initializers"
	"github.com/kulinhdev/serentyspringsedu/models"
)

func main() {
	// Load Config
	initializers.Initialize()

	// Init DB
	models.Initialize()

	// Setup API routes
	api.Initialize()

}
