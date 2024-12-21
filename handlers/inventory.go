package handlers

import (
	"net/http"
	"smart-inventory/models"
	"smart-inventory/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if result := db.Find(&products); result.Error != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve products")
			return
		}
		utils.RespondWithJSON(c, http.StatusOK, products)
	}
}

func AddProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if result := db.Create(&product); result.Error != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to add product")
			return
		}
		utils.RespondWithJSON(c, http.StatusCreated, product)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		if result := db.First(&product, id); result.Error != nil {
			utils.RespondWithError(c, http.StatusNotFound, "Product not found")
			return
		}
		if err := c.ShouldBindJSON(&product); err != nil {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}
		db.Save(&product)
		utils.RespondWithJSON(c, http.StatusOK, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if result := db.Delete(&models.Product{}, id); result.Error != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete product")
			return
		}
		utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}
