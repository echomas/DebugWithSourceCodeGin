package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UriPerson struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

var isUUID validator.Func = func(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	if _, err := uuid.Parse(fieldValue); err != nil {
		return false
	}
	return true
}

func main() {
	route := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uuid", isUUID)
	}
	route.GET("/:name/:id", func(c *gin.Context) {
		var person UriPerson
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8088")
}
