package httprsp

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Udb *sql.DB

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func test(c *gin.Context) {
	sendMail()
	c.JSON(http.StatusOK, gin.H{})
}

func loginSuc(c *gin.Context) {
	var json loginInfo
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbPass string
	var ret loginRes
	err := Udb.QueryRow("select password, auth from users where id=?", json.ID).Scan(&dbPass, &ret.Auth)
	if err != nil {
		log.Fatal(err)
	}

	ret.Status = (dbPass != json.Password)
	c.JSON(http.StatusOK, ret)
}

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	var err error
	var res bool
	if tid != "" {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM users where id=?)", tid).Scan(&res)
	} else {
		err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM users where email=?)", temail).Scan(&res)
	}
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": res})
}

func regiComplete(c *gin.Context) {
	var json regiInfo
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := Udb.Exec("INSERT INTO users VALUES(?, ?, ?)", json.ID, json.Password, json.Email)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, json)
}

func getStatus(c *gin.Context) {
	id := paramInfo{"id", 1, c.Query("id")}
	prob := paramInfo{"prob_no", 0, c.DefaultQuery("prob_no", "-1")}
	res := paramInfo{"result", 0, c.DefaultQuery("result", "-1")}
	lang := paramInfo{"lang", 0, c.DefaultQuery("lang", "-1")}
	page := c.DefaultQuery("page", "1")

	qry := makeWhere(id, prob, res, lang)
	top, _ := strconv.Atoi(page)
	top = (top - 1) * pageSize
	log.Println(qry)

	var json submitPage
	err := Udb.QueryRow("select count(*) from submits " + qry).Scan(&json.DataNum)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := Udb.Query("select * from submits "+qry+"order by subm_no desc limit ?, ?", top, pageSize)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tmp submitInfo
	for rows.Next() {
		err := rows.Scan(&tmp.SubmNo, &tmp.ID, &tmp.ProbNo, &tmp.Result,
			&tmp.Lang, &tmp.RunTime, &tmp.Memory, &tmp.Codelen, &tmp.SubmTime)
		if err != nil {
			log.Fatal(err)
		}
		json.Datas = append(json.Datas, tmp)
	}

	c.JSON(http.StatusOK, json)
}
