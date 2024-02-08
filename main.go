package main

import (
	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.LogIn)
	r.GET("/wa", controllers.Test)
	r.POST("/LogIn", controllers.LogInPost)
	r.Run()
	// pw.FuckAround("marcosignc_21", "sopitasprecio")
}
