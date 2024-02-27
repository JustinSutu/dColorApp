package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"service/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("Web 服务器正在启动...")

	if err := r.Run(); err != nil {
		log.Fatalf("无法启动 Gin Web 服务器: %v", err)
	}

	log.Println("Web 服务器启动：localhost:8080")
}
