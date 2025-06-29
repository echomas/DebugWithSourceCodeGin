package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		userGroup := apiV1.Group("/users")
		{
			//:name 是一个路径参数
			userGroup.GET("/:name", func(c *gin.Context) {
				name := c.Param("name")
				message := fmt.Sprintf("Hello %s", name)
				c.JSON(http.StatusOK, gin.H{"message": message})
			})

			//GET /api/v1/users/:name/profile
			userGroup.GET("/:name/profile", func(c *gin.Context) {
				name := c.Param("name")
				c.JSON(http.StatusOK, gin.H{
					"name":  name,
					"email": name + "@gmail.com",
					"role":  "user",
				})
			})
		}

		//GET /api/v1/files/*filepath
		apiV1.GET("/files/*filepath", func(c *gin.Context) {
			filepath := c.Param("filepath")
			c.JSON(http.StatusOK, gin.H{
				"message":    "requesting file",
				"filepath":   filepath,
				"full_match": c.FullPath(),
			})
		})
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8080")
}
