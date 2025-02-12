package main

import (
	"crud_app/initilizers"
	"crud_app/models"
)

func init() {
	initilizers.LoadEnvVariables()
	initilizers.ConnectToDB()

}

// init() initializes environment variables and database connection
// main() performs auto-migration of the User model schema

func main() {
	initilizers.DB.AutoMigrate(&models.User{}, &models.Product{})
}
