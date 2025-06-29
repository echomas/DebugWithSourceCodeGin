package main

import "github.com/gin-gonic/gin"

type UserInfo2 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

type Order2 struct {
	OrderID   string `form:"order_id"`
	UserInfo2        //匿名嵌入，没有字段名，没有 form tag
}

func main() {
	r := gin.Default()
	r.POST("/order", func(c *gin.Context) {
		var order Order2
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
