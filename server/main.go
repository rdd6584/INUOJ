package main

import (
	"INUOJ/server/httprsp"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Static("/css", "./css")
	r.Static("/img", "./img")
	r.Static("/js", "./js")
	r.LoadHTMLGlob("public/*.html")

	var err error
	httprsp.Udb, err = sql.Open("mysql", "root:20190325@tcp(127.0.0.1:3306)/inuoj")
	if err != nil {
		log.Fatal(err)
	}
	defer httprsp.Udb.Close()
	httprsp.ResRouter(r)

	r.Run(":80")
}
