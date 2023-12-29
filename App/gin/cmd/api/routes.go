package main

import (
	"ProyectoPractica/App/gin/internal/handlers"
	"ProyectoPractica/App/gin/internal/storage"
	"ProyectoPractica/App/gin/internal/services"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Crear una instancia del repositorio de productos
    productRepo := storage.NewProductRepository()

    // Crear una instancia del servicio de productos, inyectando el repositorio
    productService := services.NewProductService(productRepo)

    // Crear los handlers, pasando el servicio
    productHandler := handlers.NewProductHandler(productService)

	api := router.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.GET("/products", productHandler.ObtenerInventario)
		api.GET("/products/low-stock",productHandler.ObtenerProductosBajoStock)
	}

	return router
}
