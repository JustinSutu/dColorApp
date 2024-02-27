package routes

import (
	"github.com/gin-gonic/gin"
	"service/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", handlers.GetUsers)
}
