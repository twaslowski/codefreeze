package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	err := r.Run()
	if err != nil {
		fmt.Printf("Got error when starting webserver: %s", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "up",
		})
	})
	return r
}
