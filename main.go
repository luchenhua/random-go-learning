package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luchenhua/random-learning-go/service"
)

func main() {

	r := gin.Default()

	r.GET("/test", service.Test)

	r.Run(":8081")
}
