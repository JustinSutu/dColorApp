package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"service/routes"
	"time"
)

func main() {
	dsn := "host=localhost user=your_user password=your_password dbname=your_dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 这里可以执行数据库操作，例如自动迁移、CRUD等

	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有域
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true // 可以根据实际需要调整条件
		},
		MaxAge: 12 * time.Hour,
	}

	router.Use(cors.New(config))

	routes.RegisterRoutes(router)

	log.Println("Web 服务器正在启动...")

	if err := router.Run(); err != nil {
		log.Fatalf("无法启动 Gin Web 服务器: %v", err)
	}

	log.Println("Web 服务器启动：localhost:8080")
}
