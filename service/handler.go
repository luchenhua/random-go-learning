package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"random-learning-go/entity"

	"github.com/gin-gonic/gin"
)

type LienApiParams struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Version   string `json:"version"`
	SystemId  string `json:"systemId"`
	EventId   string `json:"eventId"`
	EventName string `json:"eventName"`
}

type LienApiResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data   string `json:"data"`
}

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

	c.SecureJSON(res.StatusCode, res.Body)
}

func TestLienApi(c *gin.Context) {

	res := sendMsgToLien(prepareApiParam(c))

	c.SecureJSON(http.StatusOK, res)
}

func prepareApiParam(c *gin.Context) LienApiParams {

	var params LienApiParams

	return params
}

func sendMsgToLien(input LienApiParams) *LienApiResponse {

	params := url.Values{"key": {"Value"}, "id": {"123"}}
	res, _ := http.PostForm(os.Getenv("LIEN_API_ADDRESS"), params)
	body, _ := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var result = new(LienApiResponse)
	json.Unmarshal(body, &result)
	return result
}
