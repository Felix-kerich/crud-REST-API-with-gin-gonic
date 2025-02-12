package controller

import (
	"crud_app/initilizers"
	"crud_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles product creation
func CreateProduct(ctx *gin.Context) {
	// Get data from request body
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create product in database
	result := initilizers.DB.Create(&product)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	// Return the created product
	ctx.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

// GetProduct retrieves a single product by ID
func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var product models.Product
	result := initilizers.DB.Preload("User").First(&product, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

// GetProducts retrieves all products
func GetProducts(ctx *gin.Context) {
	var products []models.Product
	result := initilizers.DB.Preload("User").Find(&products)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch products",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

// UpdateProduct updates a product by ID
func UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	// Get data from request body
	var body models.Product
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the product
	var product models.Product
	result := initilizers.DB.First(&product, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Update product
	initilizers.DB.Model(&product).Updates(models.Product{
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

// DeleteProduct deletes a product by ID
func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	result := initilizers.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
