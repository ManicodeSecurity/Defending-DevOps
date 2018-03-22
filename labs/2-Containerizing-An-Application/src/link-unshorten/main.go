package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	port := ":8080"
	//port := ":" + os.Getenv("PORT")

	r := gin.Default()
	r.Use(globalRecover)

	//route groups
	health := r.Group("/healthz")
	{
		health.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "healthy",
			})
		})
	}
	rest := r.Group("/api")
	{
		rest.GET("/check", UnshortenLinkAPI)
	}
	r.Run(port)
}

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			// that recovery also handle XHR's
			if XHR(c) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": rec,
				})
			} else {
				log.Println("here!")
			}
		}
	}(c)
	c.Next()
}

func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

// isValidUrl tests a string to determine if it is a url or not.
func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}
