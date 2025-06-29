package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/json", SomeHandler2)
	/*
		router.GET("/someJSON", func(c *gin.Context) {
			data := map[string]interface{}{
				"lang": "GO 语言",
				"tag":  "<br>",
			}
			c.AsciiJSON(http.StatusOK, data)
		})

			router.SetTrustedProxies(nil)
			router.TrustedPlatform = gin.PlatformCloudflare
			router.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})

	*/
	router.Run()
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler1(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {

	}
}

func SomeHandler2(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
		c.String(http.StatusOK, `the body should be formB XML`)
	}
}
