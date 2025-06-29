package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegistration1 struct {
	Username string `json:"username" binding:"required,min=4,max=12"`
	Password string `json:"password" binding:"required,gte=6,lte=20"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	router := gin.Default()
	router.POST("/register", func(c *gin.Context) {
		var user UserRegistration1

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"message":  "用户注册成功",
			"username": user.Username,
			"email":    user.Email,
		})

	})

	router.Run(":8080")
}
