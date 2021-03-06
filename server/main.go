package main

import (
	rsp "INUOJ/server/httprsp"
	"database/sql"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	r := gin.Default()
	r.Static("/css", "./css")
	r.Static("/js", "./js")
	r.Static("/img", "./img")
	r.LoadHTMLGlob("*.html")

	var err error
	rsp.Udb, err = sql.Open("mysql", rsp.OurMysql)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Udb.Close()
	rsp.ResRouter(r)

	r.Run(":80")
}
