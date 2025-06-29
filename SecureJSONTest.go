package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		//c.SecureJSON(http.StatusOK, names)
		c.JSON(200, names)
	})
	router.Run()
}
