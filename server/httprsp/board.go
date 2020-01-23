package httprsp

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPostList(c *gin.Context) {
	page := c.Query("page")
	author := paramInfo{"pt.id", 1, c.Query("author")}
	probNo := paramInfo{"pt.prob_no", 0, c.Query("prob_no")}
	title := paramInfo{"pt.title", 1, c.Query("title")}
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

	rows, err := Udb.Query("select pt.*, ifnull(rl.result,0) as result " + qry +
		makeWhere(author, probNo, title, category))
	defer rows.Close()
	printErr(err)

	var po postInfo
	var postList, noticeList []postInfo
	var no bool
	for rows.Next() {
		err = rows.Scan(&po.PostNo, &po.Title, &po.ID, &po.Category, &no, &po.ProbNo,
			&po.CmtNo, &po.PostTime, &po.Result)
		printErr(err)
		postList = append(postList, po)
	}

	rows2, err := Udb.Query("select pt.*, ifnull(rl.result,0) as result " + qry + "where notice=1")
	printErr(err)
	defer rows2.Close()

	for rows.Next() {
		err = rows.Scan(&po.PostNo, &po.Title, &po.ID, &po.Category, &no, &po.ProbNo,
			&po.CmtNo, &po.PostTime, &po.Result)
		printErr(err)
		noticeList = append(noticeList, po)
	}
	c.JSON(http.StatusOK, gin.H{"data_num": dataNum, "notice": noticeList, "normal": postList})
}

func addNewPost(c *gin.Context) {
	var po postInfo
	if err := c.ShouldBind(&po); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	var postNo string
	tx, err := Udb.Begin()
	panicErr(err)
	_, err = tx.Exec("insert into posts(title, id, prob_no, category)", po.Title, po.ID, po.ProbNo, po.Category)
	panicErr(err)
	err = tx.QueryRow("select last_insert_id()").Scan(&postNo)
	panicErr(err)

	err = tx.Commit()
	panicErr(err)

	err = ioutil.WriteFile(postDir+postNo+"/content.txt", []byte(po.Content), 0644)
	printErr(err)
	err = ioutil.WriteFile(postDir+postNo+"/code.txt", []byte(po.Code), 0644)
	printErr(err)

	c.String(http.StatusOK, "")
}

func addNewComment(c *gin.Context) {

}
