package httprsp

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func changeStat(c *gin.Context) {
	var json modStat
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	var stat int
	err = Udb.QueryRow("select stat from probs where ori_no=?", json.OriNo).Scan(&stat)
	if err != nil {
		c.String(http.StatusNotFound, "")
		return
	}
	if stat != json.FromStat {
		c.String(http.StatusInternalServerError, "fromstat is dif")
		return
	}

	_, err = Udb.Exec("update probs set stat=? where ori_no=?", json.ToStat, json.OriNo)
	if err != nil {
		log.Println(err)
		return
	}
	if json.FromStat == 0 {
		var probNo int
		err = Udb.QueryRow("select max(prob_no) from probs").Scan(&probNo)
		printErr(err)

		_, err = Udb.Exec("update probs set prob_no=? where ori_no=?", probNo+1, json.OriNo)
		printErr(err)

		moveFile(strconv.Itoa(json.OriNo))
	}
	c.String(http.StatusOK, "")
}

func myProbList(c *gin.Context) {
	id := c.Query("id")
	rows, err := Udb.Query("select pr.ori_no, pr.prob_no, pr.title, pr.stat "+
		"from probs as pr join prob_auth as pa where pr.ori_no=pa.ori_no and pa.id=? order by pr.ori_no desc", id)
	printErr(err)
	defer rows.Close()

	var ret []probListMy
	var tmp probListMy
	for rows.Next() {
		_ = rows.Scan(&tmp.OriNo, &tmp.ProbNo, &tmp.Title, &tmp.Stat)
		ret = append(ret, tmp)
	}
	c.JSON(http.StatusOK, gin.H{"problems": ret})
}

func getProbList(c *gin.Context) {
	userID := c.Query("id")
	page := c.Query("page")
	title := c.Query("title")

	if userID == "" || page == "" {
		c.String(http.StatusBadRequest, "")
		return
	}
	top, _ := strconv.Atoi(page)
	top = (top - 1) * pageSize

	qry := "from probs as pr left join result_list as rl on pr.prob_no=rl.prob_no and rl.id=? where (pr.stat=1 or pr.stat=2) "
	if title != "" {
		qry += "and rl.title like '%" + title + "%' "
	}

	var json probListPage
	var tmp probListAll
	var err error
	err = Udb.QueryRow("select count(*) "+qry, userID).Scan(&json.DataNum)
	printErr(err)

	qry += "order by pr.prob_no desc limit ?, ?"
	rows, err := Udb.Query("select pr.title, pr.prob_no, pr.attempt, pr.accept, pr.stat, ifnull(rl.result,0) as result "+
		qry, userID, top, pageSize)
	printErr(err)
	for rows.Next() {
		err = rows.Scan(&tmp.Title, &tmp.ProbNo, &tmp.Attempt, &tmp.Accept, &tmp.Stat, &tmp.Result)
		printErr(err)
		json.Problems = append(json.Problems, tmp)
	}
	c.JSON(http.StatusOK, json)
}

func getNewOriNo(c *gin.Context) {
	var ret int
	var err error
	id := c.Query("id")
	tx, err := Udb.Begin()
	panicErr(err)
	defer func() {
		tx.Rollback()
	}()

	_, err = tx.Exec("insert into probs(owner) values(?)", id)
	panicErr(err)
	err = tx.QueryRow("select last_insert_id()").Scan(&ret)
	panicErr(err)
	_, err = tx.Exec("insert into prob_auth values(?,?)", id, ret)
	panicErr(err)
	tx.Commit()

	_ = os.MkdirAll(privDir+strconv.Itoa(ret)+inDir, os.ModePerm)
	_ = os.MkdirAll(privDir+strconv.Itoa(ret)+outDir, os.ModePerm)

	c.JSON(http.StatusOK, gin.H{"ori_no": ret})
}

func uploadDesc(c *gin.Context) {
	var pb probDetail
	var err error
	if err = c.ShouldBind(&pb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(pb.SampleIn) != len(pb.SampleOut) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "in, out dif"})
		return
	}

	_, err = Udb.Exec("update probs set t_limit=?, m_limit=?, title=? where ori_no=?",
		pb.TimeLimit, pb.MemoryLimit, pb.Title, pb.OriNo)
	printErr(err)

	dir := makeDir(pb.OriNo)
	for i := 0; i < len(descName); i++ {
		err = ioutil.WriteFile(dir+descName[i], []byte(pb.Description[i]), 0644)
		printErr(err)
	}

	initSample(dir)
	for i := len(pb.SampleIn) - 1; i >= 0; i-- {
		err = ioutil.WriteFile(dir+inDir+"sample"+strconv.Itoa(i+1)+".in", []byte(pb.SampleIn[i]), 0644)
		printErr(err)
		err = ioutil.WriteFile(dir+outDir+"sample"+strconv.Itoa(i+1)+".out", []byte(pb.SampleOut[i]), 0644)
		printErr(err)
	}
}

func uploadData(c *gin.Context) {
	var pb probData
	var err error
	if err = c.ShouldBind(&pb); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	if !isPrivate(pb.OriNo) {
		c.String(http.StatusForbidden, "")
		return
	}
	if len(pb.Input) != len(pb.Output) {
		c.String(http.StatusBadRequest, "in, out not match")
		return
	}

	dir := makeDir(pb.OriNo)
	for i := len(pb.Input) - 1; i >= 0; i-- {
		_ = c.SaveUploadedFile(pb.Input[i], dir+inDir+pb.Input[i].Filename)
		_ = c.SaveUploadedFile(pb.Output[i], dir+outDir+pb.Output[i].Filename)
	}
	c.JSON(http.StatusOK, "")
}

func discardData(c *gin.Context) {
	var files fileArray
	if err := c.ShouldBind(&files); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !isPrivate(files.OriNo) {
		c.String(http.StatusForbidden, "")
		return
	}

	dir := makeDir(files.OriNo)
	for _, file := range files.FileList {
		_ = os.Remove(dir + inDir + file + ".in")
		_ = os.Remove(dir + outDir + file + ".out")
	}
	c.String(http.StatusOK, "")
}

func viewMyProbDetail(c *gin.Context) {
	var pb probDetail
	var err error
	var pNum int
	ori := c.Param("ori_no")
	pNum, err = strconv.Atoi(ori)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dir, qry string
	qry = "select t_limit, m_limit, owner, title, stat from probs where ori_no=?"
	err = Udb.QueryRow(qry, pNum).Scan(&pb.TimeLimit, &pb.MemoryLimit, &pb.Owner, &pb.Title, &pb.Stat)
	if err != nil {
		c.String(http.StatusNotFound, "")
		return
	}

	dir = makeDir(pNum)
	files, err := ioutil.ReadDir(dir + inDir)
	printErr(err)

	for _, file := range files {
		fileName := file.Name()
		if !strings.HasPrefix(fileName, "sample") {
			pb.Datas = append(pb.Datas, fileName[0:len(fileName)-3])
		}
	}

	var file []byte
	for i := 0; i < len(descName); i++ {
		file, err = ioutil.ReadFile(dir + descName[i])
		pb.Description = append(pb.Description, string(file))
	}

	for i := 1; ; i++ {
		file, err = ioutil.ReadFile(dir + inDir + "sample" + strconv.Itoa(i) + ".in")
		if err != nil {
			break
		}
		pb.SampleIn = append(pb.SampleIn, string(file))

		file, err = ioutil.ReadFile(dir + outDir + "sample" + strconv.Itoa(i) + ".out")
		if err != nil {
			log.Println(err)
			break
		}
		pb.SampleOut = append(pb.SampleOut, string(file))
	}
	c.JSON(http.StatusOK, pb)
}

func viewProbDetail(c *gin.Context) {
	var pb probDetail
	var err error
	var pNum int
	prob := c.Param("prob_no")
	userID := c.Param("id")

	var dir, qry string
	pNum, err = strconv.Atoi(prob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Udb.QueryRow("select ori_no from probs where prob_no=?", prob).Scan(&pNum)
	if err != nil {
		c.String(http.StatusNotFound, "")
		return
	}

	Udb.QueryRow("select result from result_list where id=? and prob_no=?", userID, pNum).Scan(&pb.Result)
	qry = "select t_limit, m_limit, attempt, accept, owner, title, stat from probs where ori_no=?"
	err = Udb.QueryRow(qry, pNum).Scan(&pb.TimeLimit, &pb.MemoryLimit, &pb.Attempt, &pb.Accept,
		&pb.Owner, &pb.Title, &pb.Stat)

	dir = makeDir(pNum)
	var file []byte
	for i := 0; i < len(descName); i++ {
		file, err = ioutil.ReadFile(dir + descName[i])
		pb.Description = append(pb.Description, string(file))
	}

	for i := 1; ; i++ {
		file, err = ioutil.ReadFile(dir + inDir + "sample" + strconv.Itoa(i) + ".in")
		if err != nil {
			break
		}
		pb.SampleIn = append(pb.SampleIn, string(file))

		file, err = ioutil.ReadFile(dir + outDir + "sample" + strconv.Itoa(i) + ".out")
		if err != nil {
			log.Println(err)
			break
		}
		pb.SampleOut = append(pb.SampleOut, string(file))
	}
	c.JSON(http.StatusOK, pb)
}

// *******************************funcs
func makeDir(oriNo int) string {
	if isPrivate(oriNo) {
		return privDir + strconv.Itoa(oriNo)
	}
	return pubDir + strconv.Itoa(oriNo)
}

func initSample(dir string) {
	var err error
	for i := 1; ; i++ {
		err = os.Remove(dir + inDir + "sample" + strconv.Itoa(i) + ".in")
		if err != nil {
			break
		}
		_ = os.Remove(dir + outDir + "sample" + strconv.Itoa(i) + ".out")
	}
}

func isPrivate(oriNo int) bool {
	var stat int
	err := Udb.QueryRow("select stat from probs where ori_no=?", oriNo).Scan(&stat)
	printErr(err)
	return (stat == 0)
}

func moveFile(oriNo string) {
	err := os.Rename(privDir+oriNo, pubDir+oriNo)
	printErr(err)
}
