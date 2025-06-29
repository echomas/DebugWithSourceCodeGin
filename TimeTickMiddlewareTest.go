package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latencyTime := end.Sub(start)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()

		log.Printf("[GIN] %d | %13v | %s | %s", statusCode, latencyTime, reqMethod, reqUri)
	}
}

func main() {
	router := gin.Default()
	router.Use(RequestLogger())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
