package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person2 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage1)
	router.Run(":8085")
}

func startPage1(c *gin.Context) {
	var person1 Person2
	if c.ShouldBindQuery(&person1) == nil {
		log.Println("===== only bind by query string =====")
		log.Println(person1.Name)
		log.Println(person1.Address)
	}
	c.String(200, "Success")
}
