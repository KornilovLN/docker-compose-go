package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url": "gooogle.com",
		})
	})
	r.Run() // Запустит сервер на порту 8080 по умолчанию
}
