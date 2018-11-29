package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* HTTP 请求的返回结果 */
type Result struct {
	Data    string `json:"data"`
	Message string `json:"msg"`
}

func main() {

	r := gin.Default()

	r.GET("/test", test)

	r.Run(":8081")
}

func test(c *gin.Context) {

	var result Result

	result.Data = "Go REST service"
	result.Message = "No message"

	c.SecureJSON(http.StatusOK, result)
}
