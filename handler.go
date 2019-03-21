package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Result represents the response of the HTTP request.
type Result struct {
	Data    string `json:"data"`
	Message string `json:"msg"`
}

func test(c *gin.Context) {

	var result Result

	result.Data = "Go REST service"
	result.Message = "No message"

	c.SecureJSON(http.StatusOK, result)
}
