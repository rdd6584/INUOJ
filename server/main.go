package main

import (
	"httprsp"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/templates/static", "./templates/static")
	r.LoadHTMLGlob("templates/htmls/*")

	httprsp.ResRouter(r)

	r.Run(":80")
}
