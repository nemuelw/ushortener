package main

import (
	"fmt"
	"net/http"

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
	longURL := c.PostForm("longURL")
	id, _ := shortid.Generate()
	host := c.Request.Host

	shortURL := fmt.Sprintf("https://%s/%s", host, id)
	urlMap[id] = longURL

	c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func redirectToURL(c *gin.Context) {
	hash := c.Param("hash")
	longURL, ok := urlMap[hash]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, longURL)
}
