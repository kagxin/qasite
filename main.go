package main

import (
	models "qasite/model"

	"github.com/gin-gonic/gin"
)

func init() {
	models.DB.AutoMigrate(&models.User{})
}
func main() {
	defer models.DB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
