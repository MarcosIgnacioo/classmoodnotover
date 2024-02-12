package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob(dir + "/public/templates/*")
	r.Static("/assets", dir+"/assets")
	r.GET("/", controllers.LogIn)
	r.GET("/wa", controllers.Test)
	r.POST("/LogIn", controllers.LogInPost)
	port := os.Getenv("STATE")
	if port == "dev" {
		port = "3030"
	} else {
		port = "3000"
	}
	r.Run("0.0.0.0:" + port)
}
