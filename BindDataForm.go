package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()
	router.GET("/testing", startPage)
	router.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.String(200, "Success")
}

func formMyFormHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{
		"color": fakeForm.Colors,
	})
}
