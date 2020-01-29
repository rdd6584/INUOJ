package main

import (
	jg "INUOJ/judgeServer/judge"
	"database/sql"
	"log"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	var err error
	jg.Udb, err = sql.Open("mysql", jg.OurMysql)
	if err != nil {
		log.Panicln(err)
	}
	defer jg.Udb.Close()

	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {
		jg.ReadQueue()
	}
}
