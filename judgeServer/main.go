package main

import (
	jg "INUOJ/judgeServer/judge"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error
	jg.Udb, err = sql.Open("mysql", jg.OurMysql)
	if err != nil {
		log.Panicln(err)
	}
	defer jg.Udb.Close()

	c := time.Tick(50 * time.Millisecond)
	jg.Ch = make(chan struct{}, 0)

	for now := range c {
		fmt.Println(now)
		go jg.ReadQueue()
		<-jg.Ch
	}

}
