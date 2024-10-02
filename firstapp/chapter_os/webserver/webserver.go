package webserver

import (
	logging "firstapp/chapter_os/log"

	"github.com/gin-gonic/gin"
)

const (
	PORT = ":8081"
)

func StartWebServer() {
	logging.LogInfo("Starting web server")
	logging.LogInfo("Links: {/ping /logs}")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"url": "google.com",
		})
	})

	r.GET("/logs", func(c *gin.Context) {
		c.String(200, logging.LogBuffer.String())
	})

	r.Run(PORT)
}
