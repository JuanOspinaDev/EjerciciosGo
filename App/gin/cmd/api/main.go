package main

import (
	"fmt"
)

func main() {
	router := setupRouter()

    fmt.Println("Servidor iniciado en el puerto 8080")
	router.Run(":8080")
}
