package main

import (
	"contents-api/internal/app"
	routes "contents-api/internal/routes"
	"log"

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
	routes.AppRoutes(server, dependencies.VodContentController, dependencies.ImageController, dependencies.ContentDistributorController)

	log.Printf("Servidor rodando na http://%s%s\n", ip, port)

	server.Run(port)
}
