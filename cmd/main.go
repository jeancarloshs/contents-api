package main

import (
	"log"
	"nx-api/internal/app"
	routes "nx-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	ip := "127.0.0.1"
	port := ":8000"
	server := gin.Default()

	dependencies, err := app.InitializeDependencies()
	if err != nil {
		log.Fatalf("Erro ao inicializar dependências: %v", err)
	}
	routes.AppRoutes(server, dependencies.VodContentController, dependencies.ImageController)

	log.Printf("Servidor rodando na http://%s%s\n", ip, port)

	server.Run(port)
}
