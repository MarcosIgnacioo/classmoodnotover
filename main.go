package main

import (
	"fmt"
	"os"

	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
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
		fmt.Println("dev")
		port = "3000"
	} else {
		fmt.Println("prod")
		port = "3030"
	}
	r.Run("0.0.0.0:" + port)
}
