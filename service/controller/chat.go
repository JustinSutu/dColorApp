package controller

import (
	"net/http"

	"service/service"

	"github.com/gin-gonic/gin"
)

type ChatRequest struct {
	Question string `json:"question"`
}

func Chat(c *gin.Context) {
	// 初始化结构体用于绑定请求体数据
	var chatReq ChatRequest

	if err := c.BindJSON(&chatReq); err != nil {
		// 如果绑定出错，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	params := service.ChatInfo{
		Query: chatReq.Question,
	}

	data := service.AiChat(params)

	resp := APIResponse{
		Ret:     1,
		Status:  http.StatusOK,
		Message: "Successfully response chat",
		Data:    data,
	}

	// 示例：返回一些用户数据
	c.JSON(http.StatusOK, resp)
}
