package main

import (
	"github.com/Guizzs26/go-first-api/controller"
	"github.com/Guizzs26/go-first-api/db"
	"github.com/Guizzs26/go-first-api/repository"
	"github.com/Guizzs26/go-first-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	connection, _ := db.ConnectDB()
	productRepository := repository.NewProductRepository(connection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}
