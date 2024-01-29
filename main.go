package main

import (
	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/templates/*")
	r.LoadHTMLGlob("public/templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
