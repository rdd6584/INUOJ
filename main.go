package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*")
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "mainPage.html", gin.H{})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})

	r.POST("/regi-complete", func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:20190325@tcp(127.0.0.1:3306)/inuoj")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		id := c.PostForm("id")
		pass := c.PostForm("pass")
		email := c.PostForm("email")

		result, err := db.Exec("INSERT INTO users VALUES(?, ?, ?)", id, pass, email)
		if err != nil {
			log.Fatal(err)
		}
		n, err := result.RowsAffected()
		if n == 1 {
			fmt.Println("1 row inserted.")
		}
		c.HTML(http.StatusOK, "regi-complete.html", gin.H{})
	})

	r.Run(":80")
}
