package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")
		//请求前
		c.Next()
		//请求后
		latency := time.Since(t)
		log.Print(latency)
		//获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	router.Run(":8080")
}
