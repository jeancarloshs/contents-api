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

	deps, err := app.InitializeDependencies()
	if err != nil {
		log.Fatalf("Erro ao inicializar dependências: %v", err)
	}
	routes.LoadControllers(deps)
	routes.AppRoutes(server)

	log.Printf("Servidor rodando na http://%s%s\n", ip, port)

	server.Run(port)
}
