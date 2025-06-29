package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() //step1: Process the request first.

		//step2: Check if any errors were added to the context.
		if len(c.Errors) > 0 {
			//step3: Use the last error
			err := c.Errors.Last().Err
			//step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err.Error(),
			})
		}

		//Any other steps if no errors are found
	}
}

func main() {
	router := gin.Default()

	router.Use(ErrorHandler())

	router.GET("/ok", func(c *gin.Context) {
		somethingWentWrong := false

		if somethingWentWrong {
			c.Error(errors.New("something went wrong"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	router.GET("/error", func(c *gin.Context) {
		somethingWentWrong := true

		if somethingWentWrong {
			c.Error(errors.New("something went wrong"))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Everything is fine!",
		})
	})

	router.Run()

}
