package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("Hello World!")
	})
	engine.POST("/hello", func(context *gin.Context) {
		fmt.Println("Hello World!2")
	})
	engine.Run(":9091")
}
