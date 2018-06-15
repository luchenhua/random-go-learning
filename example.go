package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/test", test)

	r.Run(":8081")
}

func test(c *gin.Context) {
	var result struct {
		Data    string `json:"data"`
		Message string `json:"msg"`
	}

	result.Data = "Go REST service"
	result.Message = "No message"

	c.JSON(http.StatusOK, result)
}
