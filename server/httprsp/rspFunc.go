package httprsp

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Udb *sql.DB

func getUserCode(c *gin.Context) {
	submNo := c.Param("subm_no")
	var subm submitInfo
	err := Udb.QueryRow("select * from submits where subm_no=?", submNo).Scan(&subm.SubmNo,
		&subm.ID, &subm.ProbNo, &subm.Result, &subm.Lang, &subm.RunTime, &subm.Memory, &subm.Codelen, &subm.SubmTime)
	printErr(err)

	err = Udb.QueryRow("select title from probs where prob_no=?", subm.ProbNo).Scan(&subm.ProbTitle)
	printErr(err)

	dir := codeDir + submNo
	code, err := ioutil.ReadFile(dir + fileType(subm.Lang))
	printErr(err)
	msg, err := ioutil.ReadFile(dir + ".txt")
	printErr(err)

	c.JSON(http.StatusOK, gin.H{"code": string(code), "compile_msg": string(msg), "submit": subm})
}

func probSubmit(c *gin.Context) {
	var json newSubmit
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx, err := Udb.Begin()
	panicErr(err)
	defer func() {
		recover()
		tx.Rollback()
	}()

	code := []byte(json.UserCode)
	_, err = tx.Exec("insert into submits(id, prob_no, lang, codelen) values(?,?,?,?)",
		json.ID, json.ProbNo, json.Lang, len(code))
	panicErr(err)

	var res, oriNo int
	err = tx.QueryRow("select last_insert_id()").Scan(&res)
	panicErr(err)

	err = tx.QueryRow("select ori_no from probs where prob_no=?", json.ProbNo).Scan(&oriNo)
	panicErr(err)

	_, err = tx.Exec("insert into judge_q values(?,?,?)", res, oriNo, json.Lang)
	panicErr(err)

	err = tx.Commit()
	panicErr(err)

	err = ioutil.WriteFile(codeDir+strconv.Itoa(res)+fileType(json.Lang), code, 0644)
	printErr(err)

	c.String(http.StatusOK, "")
}

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	var err error
	var res bool
	if tid != "" {
		err = Udb.QueryRow("select not exists (select * from users where id=?)", tid).Scan(&res)
	} else {
		err = Udb.QueryRow("select not exists (select * from users where email=?)", temail).Scan(&res)
	}
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": res})
}

func regiComplete(c *gin.Context) {
	var json regiInfo
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9\\_]+@inu\\.ac\\.kr", json.Email); !match {
		c.JSON(http.StatusBadRequest, gin.H{"status": "noINU"})
		return
	}

	var res bool
	err = Udb.QueryRow("select not exists (select * from users where id=? or email=?)", json.ID, json.Email).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	if res {
		_, err = Udb.Exec("insert into users(id, password, email) values(?, ?, ?)", json.ID, json.Password, json.Email)
		if err != nil {
			log.Println(err)
		}
		go sendMail(json.Email)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

func reSendMail(c *gin.Context) {
	var json user
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res bool
	err = Udb.QueryRow("select auth from users where id=?", json.ID).Scan(&res)
	if err != nil {
		log.Println(err)
	}
	if !res {
		var email string
		err = Udb.QueryRow("select email from users where id=?", json.ID).Scan(&email)
		if err != nil {
			log.Println(err)
		}
		var cnt int
		err = Udb.QueryRow("select count from authtokens where email=?", email).Scan(&cnt)
		if err != nil {
			log.Println(err)
		}

		if cnt >= 5 {
			c.JSON(http.StatusOK, gin.H{"status": "fail"})
		} else {
			go sendMail(email)
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "done"})
	}
}

func getStatus(c *gin.Context) {
	id := paramInfo{"id", 1, c.Query("id")}
	prob := paramInfo{"prob_no", 0, c.Query("prob_no")}
	res := paramInfo{"result", 0, c.Query("result")}
	lang := paramInfo{"lang", 0, c.Query("lang")}
	page := c.Query("page")

	qry := makeWhere(id, prob, res, lang)
	top, _ := strconv.Atoi(page)
	top = (top - 1) * pageSize

	var json submitPage
	err := Udb.QueryRow("select count(*) from submits " + qry).Scan(&json.DataNum)
	printErr(err)

	rows, err := Udb.Query("select * from submits "+qry+"order by subm_no desc limit ?, ?", top, pageSize)
	printErr(err)
	defer rows.Close()

	var tmp submitInfo
	for rows.Next() {
		err := rows.Scan(&tmp.SubmNo, &tmp.ID, &tmp.ProbNo, &tmp.Result,
			&tmp.Lang, &tmp.RunTime, &tmp.Memory, &tmp.Codelen, &tmp.SubmTime)
		printErr(err)
		err = Udb.QueryRow("select title from probs where prob_no=?", tmp.ProbNo).Scan(&tmp.ProbTitle)
		printErr(err)

		json.Datas = append(json.Datas, tmp)
	}

	c.JSON(http.StatusOK, json)
}

func emailAuth(c *gin.Context) {
	var json authInfo
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res bool
	tx, err := Udb.Begin()
	panicErr(err)
	defer tx.Rollback()
	// todo : column 한 줄 체크 필요
	tx.QueryRow("select exists (select * from authtokens where email=? and token=?)", json.Email, json.Token).Scan(&res)
	if res {
		var userID string
		err = tx.QueryRow("select id from users where email=?", json.Email).Scan(&userID)
		panicErr(err)
		_, err = tx.Exec("update users set auth=1 where id=?", userID)
		panicErr(err)
		_, err = tx.Exec("delete from authtokens where email=?", json.Email)
		panicErr(err)
		_, err = tx.Exec("insert into user_info(id) values(?)", userID)
		panicErr(err)
		err = tx.Commit()
		panicErr(err)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
