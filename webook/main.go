package main

import (
	"ast/webook/internal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	handler := web.NewUserHandler()
	handler.RegisterRouter(server.Group("user"))
	server.Run(":8080")
}
