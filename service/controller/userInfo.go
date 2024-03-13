package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func GetUserInfo(c *gin.Context) {
	// 从路径中获取userId
	userId := c.Param("userId")

	// 初始化结构体用于绑定请求体数据
	var userReq UserRequest
	if err := c.BindJSON(&userReq); err != nil {
		// 如果绑定出错，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 构造一个简单的用户信息响应，实际情况中可能需要根据参数从数据库获取信息
	userInfo := map[string]interface{}{
		"userId": userId,
		"age":    userReq.Age,
		"name":   userReq.Name,
	}

	resp := APIResponse{
		Ret:     1,
		Status:  http.StatusOK,
		Message: "Successfully retrieved users.",
		Data:    userInfo,
	}

	// 示例：返回一些用户数据
	c.JSON(http.StatusOK, resp)
}
