package main

import (
	"GinDemo1/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/static", "./static")
	r.GET("/userLogi", controller.UserLogin)
	r.Run(":8080")
}
