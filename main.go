package main

import (
	"github.com/MarcosIgnacioo/classmoodls/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func main() {
	state := os.Getenv("STATE")
	var dir string
	var err error
	if state == "prod" {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		dir = "."
	}
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob(dir + "/public/templates/*")
	r.Static("/assets", dir+"/assets")
	r.GET("/", controllers.LogIn)
	r.GET("/wa", controllers.Test)
	r.POST("/LogIn", controllers.LogInPost)
	var port string
	if state == "dev" {
		port = "3030"
	} else {
		port = "3000"
	}
	r.Run("0.0.0.0:" + port)
}
