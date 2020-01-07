package httprsp

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var Udb *sql.DB

func editUserInfo(c *gin.Context) {
	var json editInfo
	var err error
	if err = c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(json)
	if !isCorrectInfo(json.ID, json.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		return
	}
	editPass(json.ID, json.NewPassword)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func regiValid(c *gin.Context) {
	tid := c.Query("id")
	temail := c.Query("email")
	var err error
	var res bool
	log.Println("pass")
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
	if match, _ := regexp.MatchString("^[a-zA-Z0-9]+@inu\\.ac\\.kr", json.Email); !match {
		c.JSON(http.StatusBadRequest, gin.H{"status": "noINU"})
		return
	}

	var res bool
	err = Udb.QueryRow("select not exists (select * from users where id=? or email=?)", json.ID, json.Email).Scan(&res)
	if err != nil {
		log.Fatal(err)
	}
	if res {
		_, err = Udb.Exec("insert into users values(?, ?, ?, 0)", json.ID, json.Password, json.Email)
		if err != nil {
			log.Println(err)
		}
		go sendMail(json.Email)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

func reSendMail(c *gin.Context) { // 횟수 제한 구현 필요
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
	log.Println(id)
	prob := paramInfo{"prob_no", 0, c.Query("prob_no")}
	log.Println(prob)
	res := paramInfo{"result", 0, c.Query("result")}
	log.Println(res)
	lang := paramInfo{"lang", 0, c.Query("lang")}
	log.Println(lang)
	page := c.Query("page")
	log.Println(page)

	qry := makeWhere(id, prob, res, lang)
	top, _ := strconv.Atoi(page)
	top = (top - 1) * pageSize

	var json submitPage
	err := Udb.QueryRow("select count(*) from submits " + qry).Scan(&json.DataNum)
	if err != nil {
		log.Println(err)
	}

	rows, err := Udb.Query("select * from submits "+qry+"order by subm_no desc limit ?, ?", top, pageSize)
	if err != nil {
		log.Println(err)
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

func emailAuth(c *gin.Context) {
	var json authInfo
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var res bool
	tx, err := Udb.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()
	// todo : column 한 줄 체크 필요
	tx.QueryRow("select exists (select * from authtokens where email=? and token=?)", json.Email, json.Token).Scan(&res)
	if res {
		_, err = tx.Exec("update users set auth=1 where email=?", json.Email)
		if err != nil {
			log.Panic(err)
		}
		_, err = tx.Exec("delete from authtokens where email=?", json.Email)
		if err != nil {
			log.Panic(err)
		}
		err = tx.Commit()
		if err != nil {
			log.Panic(err)
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "fail"})
	}
}

func toMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
