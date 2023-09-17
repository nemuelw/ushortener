package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/:hash", controllers.GetURL)

	r.Run("localhost:1234")
}
