package service

import (
	"net/http"
	"random-learning-go/entity"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	var result entity.Result

	result.Data = "Go REST service"
	result.Message = "No message."

	c.SecureJSON(http.StatusOK, result)
}
