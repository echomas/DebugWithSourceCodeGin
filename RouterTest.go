package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	{
		v1 := router.Group("/v1")
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	{
		v2 := router.Group("/v2")
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}
}

func loginEndpoint(c *gin.Context) {
	log.Println("loginEndpoint")
}

func submitEndpoint(c *gin.Context) {
	log.Println("submitEndpoint")
}

func readEndpoint(c *gin.Context) {
	log.Println("readEndpoint")
}
