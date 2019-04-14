package service

import (
	"net/http"
	"net/url"
	"io/ioutil"

	"fmt"
	"random-learning-go/entity"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	var result entity.Result

	result.Data = "Go REST service"
	result.Message = "No message."

	c.SecureJSON(http.StatusOK, result)
}

func TestLien(c *gin.Context) {

	params := url.Values{"key": {"Value"}, "id": {"123"}}
	res, _:= http.PostForm("http://baidu.com", params)
	defer res.Body.Close()
	body, _:= ioutil.ReadAll(res.Body)
   
	fmt.Println(string(body))
}