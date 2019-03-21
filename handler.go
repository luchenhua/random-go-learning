package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {

	var result Result

	result.Data = "Go REST service"
	result.Message = "No message"

	c.SecureJSON(http.StatusOK, result)
}
