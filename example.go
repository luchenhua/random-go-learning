package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET(
		"/test",
		func(c *gin.Context) {
			var result struct {
				Data    string `json:"data"`
				Message string `json:"msg"`
			}

			result.Data = "Go REST service"
			result.Message = "No message"

			c.JSON(http.StatusOK, result)
		})

	r.Run(":8081")
}
