package main

import (
	"crud_app/controller"
	"crud_app/initilizers"

	"github.com/gin-gonic/gin"
)

func init() {
	initilizers.LoadEnvVariables()
	initilizers.ConnectToDB()

}

func main() {
	router := gin.Default()

	// User routes
	router.POST("/users", controller.CreateUser)
	router.GET("/users", controller.GetUsers)
	router.GET("/users/:id", controller.GetUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)

	// Product routes
	router.POST("/products", controller.CreateProduct)
	router.GET("/products", controller.GetProducts)
	router.GET("/products/:id", controller.GetProduct)
	router.PUT("/products/:id", controller.UpdateProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)

	router.Run()
}
