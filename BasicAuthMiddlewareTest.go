package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "123444"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "123455"},
}

func main() {
	router := gin.Default()

	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "123",
		"lena":   "abc",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(context1 *gin.Context) {
		user := context1.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context1.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context1.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	router.Run()
}
