package service

import (
	"github.com/gin-gonic/gin"
	"github.com/luchenhua/random-learning-go/entity"
	"net/http"
)

func Test(c *gin.Context) {

	var result entity.Result

	result.Data = "Go REST service"
	result.Message = "No message"

	c.SecureJSON(http.StatusOK, result)
}
