package main

import (
	"go-restapi/initializers"
	"go-restapi/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnection()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
