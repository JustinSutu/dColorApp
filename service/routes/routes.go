package routes

import (
	controller "service/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/list", controller.GetUsers)
	router.POST("/userInfo/:userId", controller.GetUserInfo)
	router.POST("/chat", controller.Chat)
}
