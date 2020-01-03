package httprsp

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Udb *sql.DB

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
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

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	var err error
	var res string
	if tid != "" {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM users where id=?)", tid).Scan(&res)
	} else {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM usjers where email=?)", temail).Scan(&res)
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

func getStatus(c *gin.Context) { // 페이지 번호 필요
	id := paramInfo{"id", 1, c.Query("id")}
	prob := paramInfo{"prob_no", 0, c.DefaultQuery("prob_no", "-1")}
	res := paramInfo{"result", 0, c.DefaultQuery("result", "-1")}
	lang := paramInfo{"lang", 0, c.DefaultQuery("lang", "-1")}

	qry := "select * from submits " + makeWhere(id, prob, res, lang)
	qry += "order by subm_no desc limit ?, 15"
	Udb.Query(qry)
}
