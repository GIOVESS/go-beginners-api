package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Hello Giovanni endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, I'm Giovanni!",
		})
	})

	var err = r.Run()
	if err != nil {
		return
	} // Runs on localhost:8080
}
