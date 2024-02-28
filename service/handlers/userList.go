package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Ret     int         `json:"ret"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Name: "Alice", Email: "alice@example.com"},
	{Name: "Bob", Email: "bob@example.com"},
}

func GetUsers(c *gin.Context) {
	resp := APIResponse{
		Ret:     1,
		Status:  http.StatusOK,
		Message: "Successfully retrieved users.",
		Data:    users,
	}

	// 示例：返回一些用户数据
	c.JSON(http.StatusOK, resp)
}
