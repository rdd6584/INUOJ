package main

import (
	rsp "INUOJ/server/httprsp"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/img", "./img")
	r.LoadHTMLGlob("*.html")

	var err error
	rsp.Udb, err = sql.Open("mysql", "root:20190325@tcp(127.0.0.1:3306)/inuoj")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Udb.Close()
	rsp.ResRouter(r)

	r.Run(":8000")
}
