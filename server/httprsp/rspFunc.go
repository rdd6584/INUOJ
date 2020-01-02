package httprsp

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Udb *sql.DB

type regiInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}
type loginInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
}

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "mainPage.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func loginSuc(c *gin.Context) {
	var json loginInfo
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbPass string
	err := Udb.QueryRow("select password from users where id=?", json.ID).Scan(&dbPass)
	if err != nil {
		log.Fatal(err)
	}
	if dbPass != json.Password {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	var err error
	var res string
	if tid != "" {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM users where id=?)", tid).Scan(&res)
	} else {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM users where email=?)", temail).Scan(&res)
	}
	if err != nil {
		log.Fatal(err)
	}
	if res == "1" {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

func regiComplete(c *gin.Context) {
	var json regiInfo
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := Udb.Exec("INSERT INTO users VALUES(?, ?, ?)", json.ID, json.Password, json.Email)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	c.JSON(http.StatusOK, json)
}
