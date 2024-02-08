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
	ScrappedInfo, err := pw.FuckAround(user, password)
	if err != nil {
		// ver como puedo poner StatusUnauthorized
		c.HTML(http.StatusOK, "index.html", err)
		return
	}
	c.HTML(http.StatusOK, "assigments.html", ScrappedInfo)
}

func Test(c *gin.Context) {
	pw.Test()
	c.HTML(http.StatusOK, "test.html", nil)
}
