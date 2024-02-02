package controllers

import (
	"fmt"
	"net/http"

	pw "github.com/MarcosIgnacioo/classmoodls/playwright"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	fmt.Println(c.Request.Method)
	c.HTML(http.StatusOK, "index.html", nil)
}

func LogInPost(c *gin.Context) {
	user := c.PostForm("username")
	password := c.PostForm("password")
	ScrappedInfo := pw.FuckAround(user, password)
	c.HTML(http.StatusOK, "assigments.html", ScrappedInfo)
}
func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}
