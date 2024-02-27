package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	// 示例：返回一些用户数据
	c.JSON(http.StatusOK, gin.H{"message": "获取用户列表"})
}
