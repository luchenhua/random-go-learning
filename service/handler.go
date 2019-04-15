package service

import (
	"net/http"
	"net/url"

	"random-learning-go/entity"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	var result entity.Result

	result.Data = "Go REST service"
	result.Message = "No message modified."

	c.SecureJSON(http.StatusOK, result)
}

func TestLien(c *gin.Context) {

	params := url.Values{"key": {"Value"}, "id": {"123"}}
	res, _ := http.PostForm("http://163.com", params)

	defer res.Body.Close()

	c.SecureJSON(res.StatusCode, res.Status)
}
