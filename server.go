package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

var urlMap = make(map[string]string)

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenURL)
	r.GET("/:hash", redirectToURL)

	r.Run("localhost:1234")
}

func shortenURL(c *gin.Context) {

}

func redirectToURL(c *gin.Context) {

}
