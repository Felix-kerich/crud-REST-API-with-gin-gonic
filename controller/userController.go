package controller

import (
	"crud_app/initilizers"
	"crud_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser handles user creation
func CreateUser(ctx *gin.Context) {
	// Get data from request body
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user in database
	result := initilizers.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Return the created user
	ctx.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

// GetUser retrieves a single user by ID
func GetUser(ctx *gin.Context) {
	// Get id from url
	id := ctx.Param("id")

	var user models.User
	result := initilizers.DB.First(&user, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// GetUsers retrieves all users
func GetUsers(ctx *gin.Context) {
	var users []models.User
	result := initilizers.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// UpdateUser updates a user by ID
func UpdateUser(ctx *gin.Context) {
	// Get id from url
	id := ctx.Param("id")

	// Get data from request body
	var body models.User
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the user
	var user models.User
	result := initilizers.DB.First(&user, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Update user
	initilizers.DB.Model(&user).Updates(models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// DeleteUser deletes a user by ID
func DeleteUser(ctx *gin.Context) {
	// Get id from url
	id := ctx.Param("id")

	// Delete the user
	result := initilizers.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}