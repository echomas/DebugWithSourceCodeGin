package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type Order struct {
	OrderID string   `form:"order_id"`
	User    UserInfo `form:"user"`
}

func main() {
	r := gin.Default()
	r.POST("/order", func(c *gin.Context) {
		var order Order
		if err := c.ShouldBind(&order); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, order)
	})
	r.Run()
}
