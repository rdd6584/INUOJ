package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/*")
	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mainPage.html", gin.H{})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.Run(":6789")
}
