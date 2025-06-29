package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := router.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func benchEndpoint(c *gin.Context) {}

func analyticsEndpoint(c *gin.Context) {}
