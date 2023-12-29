package services

import (
    "ProyectoPractica/App/gin/internal/models"
    "ProyectoPractica/App/gin/internal/storage"
)

// ProductService define las operaciones disponibles para los productos.
type ProductService interface {
    ObtenerInventario() ([]models.Product, error)
    ObtenerProductosBajoStock(umbral int) ([]models.Product, error)
}

type productService struct {
    productRepository storage.ProductRepository
}

// NewProductService crea una nueva instancia de ProductService.
func NewProductService(repo storage.ProductRepository) ProductService {
    return &productService{
        productRepository: repo,
    }
}

// ObtenerInventario devuelve todos los productos del inventario.
func (s *productService) ObtenerInventario() ([]models.Product, error) {
    return s.productRepository.ObtenerListaProductos(), nil
}

// ObtenerProductosBajoStock devuelve los productos con cantidad menor o igual al umbral.
func (s *productService) ObtenerProductosBajoStock(umbral int) ([]models.Product, error) {
    return s.productRepository.ObtenerProductosBajoStock(umbral), nil
}
