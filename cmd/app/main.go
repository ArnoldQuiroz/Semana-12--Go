package main

import (
	"fmt"
	"hello_world/internal/usuario"
	"hello_world/pkg/logger"
)

func main() {
	fmt.Println("Iniciando aplicación...")
	logger.Info("aplicación iniciada")
	usuario.Ejecutar()
}
