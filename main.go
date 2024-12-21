package main

import (
	"log"
	"smart-inventory/handlers"
	"smart-inventory/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := services.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	router := gin.Default()

	router.GET("/products", handlers.GetAllProducts(db))
	router.POST("/products", handlers.AddProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	log.Println("Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
