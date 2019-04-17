package main

import (
	"random-learning-go/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/test", service.Test)
	r.GET("/testLien", service.TestLien)
	r.GET("/testLienApi", service.TestLienApi)

	r.Run(":3000")
}
