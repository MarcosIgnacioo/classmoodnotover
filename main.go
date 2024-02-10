package main

import (
	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/templates/*")
	r.Static("/assets", "/assets")
	r.GET("/", controllers.LogIn)
	r.GET("/wa", controllers.Test)
	r.POST("/LogIn", controllers.LogInPost)
	port := os.Getenv("STATE")
	if port == "dev" {
		port = "3000"
	} else {
		gin.SetMode(gin.ReleaseMode)
		port = "3030"
	}
	r.Run("0.0.0.0:" + port)
	// pw.FuckAround("marcosignc_21", "sopitasprecio")
}
