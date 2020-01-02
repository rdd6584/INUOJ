package httprsp

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

var Udb *sql.DB

func Init(e *gin.Engine) {
	Udb, err := sql.Open("mysql", "root:20190325@tcp(127.0.0.1:3306)/inuoj")
	if err != nil {
		log.Fatal(err)
	}
	defer Udb.Close()
}
func ResRouter(e *gin.Engine) {
	app := e.Group("/")
	{
		app.GET("/", toMain)
		app.GET("/login", login)
		app.GET("/register", register)
		app.POST("/regi-done", regiComplete)
		app.GET("/register/vaild", regiValid)
	}
}
