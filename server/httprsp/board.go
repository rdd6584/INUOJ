package httprsp

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getPostList(c *gin.Context) {
	page := c.Query("page")
	author := paramInfo{"pt.id", 1, c.DefaultQuery("author", "")}
	probNo := paramInfo{"pt.prob_no", 0, c.DefaultQuery("prob_no", "0")}
	title := paramInfo{"pt.title", 1, c.DefaultQuery("title", "")}
	category := paramInfo{"pt.category", 0, c.Query("category")}
	userID := c.Query("id")

	top, err := strconv.Atoi(page)
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	top = (top - 1) * pageSize

	qry := "from posts as pt left join result_list as rl on pt.prob_no=rl.prob_no and rl.id=? "
	var dataNum int
	err = Udb.QueryRow("select count(*) "+qry, userID).Scan(&dataNum)
	printErr(err)

	var po postInfo
	var postList, noticeList []postInfo
	var no bool

	rows2, err := Udb.Query("select pt.*, ifnull(rl.result,0) as result "+qry+"where notice=1", userID)
	printErr(err)
	defer rows2.Close()

	for rows2.Next() {
		err = rows2.Scan(&po.PostNo, &po.Title, &po.ID, &po.Category, &no, &po.ProbNo,
			&po.CmtNo, &po.PostTime, &po.Result)
		printErr(err)
		if po.ProbNo != 0 {
			err = Udb.QueryRow("select title from probs where prob_no=?", po.ProbNo).Scan(&po.ProbTitle)
			printErr(err)
		}
		noticeList = append(noticeList, po)
	}

	qry += makeWhere(author, probNo, title, category) + "order by post_no desc limit ?, ?"
	rows, err := Udb.Query("select pt.*, ifnull(rl.result,0) as result "+qry, userID, top, pageSize)
	defer rows.Close()
	printErr(err)

	for rows.Next() {
		err = rows.Scan(&po.PostNo, &po.Title, &po.ID, &po.Category, &no, &po.ProbNo,
			&po.CmtNo, &po.PostTime, &po.Result)
		printErr(err)
		if po.ProbNo != 0 {
			err = Udb.QueryRow("select title from probs where prob_no=?", po.ProbNo).Scan(&po.ProbTitle)
			printErr(err)
		} else {
			po.ProbTitle = ""
		}
		postList = append(postList, po)
	}

	c.JSON(http.StatusOK, gin.H{"data_num": dataNum, "notice": noticeList, "normal": postList})
}

func addNewPost(c *gin.Context) {
	var po postInfo
	var err error
	if err = c.ShouldBind(&po); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var postNo string
	tx, err := Udb.Begin()
	defer func() {
		tx.Rollback()
		_ = recover()
	}()
	panicErr(err)
	_, err = tx.Exec("insert into posts(title, id, prob_no, category) values(?,?,?,?)", po.Title, po.ID, po.ProbNo, po.Category)
	panicErr(err)
	err = tx.QueryRow("select last_insert_id()").Scan(&postNo)
	panicErr(err)

	err = tx.Commit()
	panicErr(err)

	err = os.MkdirAll(postDir+postNo, os.ModePerm)
	printErr(err)
	err = ioutil.WriteFile(postDir+postNo+"/content.txt", []byte(po.Content), 0644)
	printErr(err)
	err = ioutil.WriteFile(postDir+postNo+"/code.txt", []byte(po.Code), 0644)
	printErr(err)

	c.String(http.StatusOK, "")
}

func addNewComment(c *gin.Context) {
	var cmt cmtInfo
	var err error
	if err = c.ShouldBind(&cmt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var cmtNo string
	_, err = Udb.Exec("update posts set cmt_no=cmt_no+1 where post_no=?", cmt.PostNo)
	printErr(err)
	err = Udb.QueryRow("select cmt_no from posts where post_no=?", cmt.PostNo).Scan(&cmtNo)
	printErr(err)

	err = ioutil.WriteFile(postDir+strconv.Itoa(cmt.PostNo)+"/"+cmtNo+"_"+cmt.ID+".txt", []byte(cmt.Comment), 0644)
	printErr(err)
}

func viewPost(c *gin.Context) {
	postNo := c.Param("post_no")
	var po postInfo
	var no bool
	var err error
	err = Udb.QueryRow("select * from posts where post_no=?", postNo).Scan(&po.PostNo, &po.Title,
		&po.ID, &po.Category, &no, &po.ProbNo, &po.CmtNo, &po.PostTime)
	printErr(err)
	if po.ProbNo != 0 {
		err = Udb.QueryRow("select title from probs where prob_no=?", po.ProbNo).Scan(&po.ProbTitle)
		printErr(err)
	}
	content, err := ioutil.ReadFile(postDir + postNo + "/content.txt")
	printErr(err)
	po.Content = string(content)

	code, err := ioutil.ReadFile(postDir + postNo + "/code.txt")
	printErr(err)
	po.Code = string(code)

	files, err := ioutil.ReadDir(postDir + postNo)
	printErr(err)

	var cmt cmtInfo
	for _, file := range files {
		fileName := file.Name()
		if match, _ := regexp.MatchString("^[1-9]", fileName); match {
			strArr := strings.Split(fileName, "_")
			cmt.ID = strArr[1]
			cmt.CmtTime = file.ModTime().Format("2006-01-02 15:04:05")

			cmtContent, err := ioutil.ReadFile(postDir + postNo + fileName)
			printErr(err)
			cmt.Comment = string(cmtContent)
			po.CmtList = append(po.CmtList, cmt)
		} else {
			break
		}
	}
	c.JSON(http.StatusOK, postNo)
}

func setNotice(c *gin.Context) {
	var no noticeInfo
	if err := c.ShouldBind(&no); err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	_, err := Udb.Exec("update posts set notice=? where post_no=?", no.Notice, no.PostNo)
	printErr(err)
	c.String(http.StatusOK, "")
}
