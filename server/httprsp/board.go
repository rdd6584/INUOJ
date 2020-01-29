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
	category := paramInfo{"pt.category", 0, c.DefaultQuery("category", "0")}
	title := c.DefaultQuery("title", "")
	userID := c.Query("id")

	var top int
	var err error
	if top, err = strconv.Atoi(page); err != nil {
		top = 1
	}
	if _, err = strconv.Atoi(probNo.Value); err != nil {
		probNo.Value = "0"
	}
	if tmp, err := strconv.Atoi(category.Value); err != nil || tmp < 0 || tmp >= CategorySize {
		category.Value = "0"
	}
	top = (top - 1) * pageSize

	qry := "from posts as pt left join result_list as rl on pt.prob_no=rl.prob_no and rl.id=? "
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
		} else {
			po.ProbTitle = ""
		}
		noticeList = append(noticeList, po)
	}

	qry += makeWhere(author, probNo, category)
	if title != "" {
		if author.Value == "" && probNo.Value == "0" && category.Value == "0" {
			qry += "where "
		} else {
			qry += "and "
		}
		qry += "title like '%" + title + "%' "
	}

	var dataNum int
	_ = Udb.QueryRow("select count(*) "+qry, userID).Scan(&dataNum)

	rows, err := Udb.Query("select pt.*, ifnull(rl.result,0) as result "+qry+
		"order by pt.post_no desc limit ?, ?", userID, top, pageSize)
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
	result, err := Udb.Exec("insert into posts(title, id, prob_no, category) values(?,?,?,?)", po.Title, po.ID, po.ProbNo, po.Category)
	printErr(err)
	num, err := result.LastInsertId()
	printErr(err)
	postNo = strconv.Itoa(int(num))

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
	result, _ := Udb.Exec("update posts set cmt_no=cmt_no+1 where post_no=?", cmt.PostNo)
	if !affectedOneRow(result) {
		c.String(http.StatusNotFound, "post_no not exist")
		return
	}
	_ = Udb.QueryRow("select cmt_no from posts where post_no=?", cmt.PostNo).Scan(&cmtNo)

	err = ioutil.WriteFile(postDir+paddingZero(cmt.PostNo, 1000)+"/"+cmtNo+"."+cmt.ID+".txt", []byte(cmt.Comment), 0644)
	printErr(err)
}

func viewPost(c *gin.Context) {
	postNo := c.Param("post_no")
	if _, err := strconv.Atoi(postNo); err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	var po postInfo
	var no bool
	var err error
	err = Udb.QueryRow("select * from posts where post_no=?", postNo).Scan(&po.PostNo, &po.Title,
		&po.ID, &po.Category, &no, &po.ProbNo, &po.CmtNo, &po.PostTime)
	if err != nil {
		c.String(http.StatusNotFound, "post_no not exist")
		return
	}

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
			strArr := strings.Split(fileName, ".")
			cmt.ID = strArr[1]
			cmt.CmtTime = file.ModTime().Format("2006-01-02 15:04:05")

			cmtContent, err := ioutil.ReadFile(postDir + postNo + "/" + fileName)
			printErr(err)
			cmt.Comment = string(cmtContent)
			po.CmtList = append(po.CmtList, cmt)
		} else {
			break
		}
	}
	c.JSON(http.StatusOK, po)
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

func paddingZero(num int, fm int) string {
	ret := ""
	for fm > 0 {
		ret += strconv.Itoa(num / fm)
		fm /= 10
	}
	return ret
}
