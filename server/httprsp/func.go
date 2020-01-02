package httprsp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type regiInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "mainPage.html", gin.H{})
}

func login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	if tid != "" {
		result, err := Udb.Exec("SELECT EXISTS (SELECT * FROM users WHERE id=?) AS SUCCESS", tid)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"idUnique": (result == "0")})
	} else {
		c.JSON(http.StatusOK, gin.H{"idUnique": true})
	}
}

func regiComplete(c *gin.Context) {
	var json regiInfo
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*id := c.PostForm("id")
	pass := c.PostForm("pass")
	email := c.PostForm("email")

	result, err := db.Exec("INSERT INTO users VALUES(?, ?, ?)", id, pass, email)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}*/
	result, err := Udb.Exec("INSERT INTO users VALUES(?, ?, ?)", json.ID, json.Password, json.Email)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	c.JSON(http.StatusOK, json)
	//c.HTML(http.StatusOK, "regi-complete.html", gin.H{})
}
