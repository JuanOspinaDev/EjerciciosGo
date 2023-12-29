package storage

import "ProyectoPractica/App/gin/internal/models"

// ProductRepository define las operaciones del repositorio de productos.
type ProductRepository interface {
    ObtenerListaProductos() []models.Product
    ObtenerProductosBajoStock(umbral int) []models.Product
}

// productRepository es una implementación concreta de ProductRepository.
type productRepository struct {
    inventario []models.Product
}

// NewProductRepository crea una nueva instancia de un repositorio de productos.
func NewProductRepository() ProductRepository {
    return &productRepository{
        inventario: []models.Product{
            {ID: "123", Name: "Teclado", Quantity: 15, Price: 19.99},
            {ID: "124", Name: "Mouse", Quantity: 7, Price: 9.99},
            {ID: "125", Name: "Monitor", Quantity: 8, Price: 99.99},
			{ID: "126", Name: "Auriculares", Quantity: 12, Price: 29.99},
			{ID: "127", Name: "Teclado Mecánico", Quantity: 5, Price: 59.99},
			{ID: "128", Name: "Mouse Inalámbrico", Quantity: 10, Price: 19.99},
			{ID: "129", Name: "Cámara Web", Quantity: 7, Price: 39.99},
			{ID: "130", Name: "Micrófono USB", Quantity: 4, Price: 69.99},
			{ID: "131", Name: "Monitor 4K", Quantity: 13, Price: 399.99},
			{ID: "132", Name: "Impresora", Quantity: 5, Price: 89.99},
			{ID: "133", Name: "Altavoces Bluetooth", Quantity: 8, Price: 49.99},
			{ID: "134", Name: "Cargador Portátil", Quantity: 15, Price: 24.99},
			{ID: "135", Name: "SSD Externo", Quantity: 6, Price: 109.99},
			{ID: "136", Name: "Tarjeta Gráfica", Quantity: 2, Price: 499.99},
			{ID: "137", Name: "RAM DDR4 8GB", Quantity: 10, Price: 79.99},
			{ID: "138", Name: "Silla de Oficina", Quantity: 4, Price: 149.99},
			{ID: "139", Name: "Escritorio Ajustable", Quantity: 3, Price: 299.99},
			{ID: "140", Name: "Mochila para Laptop", Quantity: 8, Price: 59.99},
        },
    }
}

// ObtenerListaProductos devuelve todos los productos del inventario.
func (r *productRepository) ObtenerListaProductos() []models.Product {
    return r.inventario
}

// ObtenerProductosBajoStock devuelve los productos con cantidad menor o igual al umbral.
func (r *productRepository) ObtenerProductosBajoStock(umbral int) []models.Product {
    var bajoStock []models.Product
    for _, product := range r.inventario {
        if product.Quantity <= umbral {
            bajoStock = append(bajoStock, product)
        }
    }
    return bajoStock
}
