package main

import (
	"gin_vue_study/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)
	panic(r.Run(":8080"))
	return
}
