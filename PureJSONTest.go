package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world Now!</b>",
		})
	})

	router.Run(":8080")
}
