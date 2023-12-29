package handlers

import (
	"ProyectoPractica/App/gin/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// productHandler estructura para manejar solicitudes relacionadas con productos.
type productHandler struct {
	productService services.ProductService
}

// NewProductHandler crea un nuevo manejador de productos.
func NewProductHandler(productService services.ProductService) *productHandler {
	return &productHandler{
		productService: productService,
	}
}

// ObtenerInventario maneja la solicitud GET para obtener el inventario completo.
func (h *productHandler) ObtenerInventario(c *gin.Context) {
	productos, err := h.productService.ObtenerInventario()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

func (h *productHandler) ObtenerProductosBajoStock(c *gin.Context) {
	umbralStr := c.DefaultQuery("threshold", "10") 

	umbral, err := strconv.Atoi(umbralStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Umbral inválido"})
		return
	}
	productos, err := h.productService.ObtenerProductosBajoStock(umbral)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}
