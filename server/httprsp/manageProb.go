package httprsp

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func changeStat(c *gin.Context) {
	var json modStat
	var err error
	if err = c.ShouldBind(&json); err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	_, err = Udb.Exec("update probs set stat=? where ori_no=?", json.ToStat, json.OriNo)
	if err != nil {
		log.Println(err)
		return
	}
	if json.FromStat == 0 {
		moveFile(strconv.Itoa(json.OriNo))
	}
	c.String(http.StatusOK, "")
}

func myProbList(c *gin.Context) {
	id := c.Query("id")
	log.Println("wrrwrwrwwr")
	rows, err := Udb.Query("select pr.ori_no, pr.prob_no, pr.title, pr.stat "+
		"from probs as pr join prob_auth as pa where pr.ori_no=pa.ori_no and pa.id=? order by pr.ori_no desc", id)
	printErr(err)
	defer rows.Close()

	var ret []probForList
	var tmp probForList
	for rows.Next() {
		_ = rows.Scan(&tmp.OriNo, &tmp.ProbNo, &tmp.Title, &tmp.Stat)
		ret = append(ret, tmp)
	}
	c.JSON(http.StatusOK, gin.H{"problems": ret})
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

	_ = os.MkdirAll(privDir+strconv.Itoa(ret)+dataDir, os.ModePerm)

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

	var dir string
	if pb.ProbNo == 0 { // 동작 확인
		dir = privDir + strconv.Itoa(pb.OriNo)
	} else {
		dir = pubDir + strconv.Itoa(pb.OriNo)
	}
	for i := 0; i < len(descName); i++ {
		err = ioutil.WriteFile(dir+descName[i], []byte(pb.Description[i]), 0644)
		printErr(err)
	}

	dir += sampleDir
	initSample(dir)
	for i := len(pb.SampleIn) - 1; i >= 0; i-- {
		err = ioutil.WriteFile(dir+strconv.Itoa(i+1)+".in", []byte(pb.SampleIn[i]), 0644)
		printErr(err)
		err = ioutil.WriteFile(dir+strconv.Itoa(i+1)+".out", []byte(pb.SampleOut[i]), 0644)
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

	dir := privDir + strconv.Itoa(pb.OriNo) + dataDir + "/"
	for i := len(pb.Input) - 1; i >= 0; i-- {
		_ = c.SaveUploadedFile(pb.Input[i], dir+pb.Input[i].Filename)
		_ = c.SaveUploadedFile(pb.Output[i], dir+pb.Output[i].Filename)
	}
	c.JSON(http.StatusOK, "")
}

func discardData(c *gin.Context) {
	var files fileArray
	if err := c.ShouldBind(&files); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dir := privDir + strconv.Itoa(files.OriNo) + dataDir + "/"
	for _, file := range files.FileList {
		_ = os.Remove(dir + file + ".in")
		_ = os.Remove(dir + file + ".out")
	}
	c.String(http.StatusOK, "")
}

func viewProbDetail(c *gin.Context) {
	var pb probDetail
	var err error
	var pNum int
	qry := "select * from probs where "
	ori := c.Param("ori_no")
	prob := c.Param("prob_no")
	if ori == "" {
		qry += "prob_no=" + prob
		pNum, err = strconv.Atoi(prob)
	} else {
		qry += "ori_no=" + ori
		pNum, err = strconv.Atoi(ori)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Udb.QueryRow(qry).Scan(&pb.OriNo, &pb.ProbNo, &pb.TimeLimit,
		&pb.MemoryLimit, &pb.Attempt, &pb.Accept, &pb.Owner, &pb.Title, &pb.Stat)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dir := privDir + strconv.Itoa(pNum)
	var file []byte
	for i := 0; i < len(descName); i++ {
		file, err = ioutil.ReadFile(dir + descName[i])
		pb.Description = append(pb.Description, string(file))
	}

	dir += sampleDir
	for i := 1; ; i++ {
		file, err = ioutil.ReadFile(dir + strconv.Itoa(i) + ".in")
		if err != nil {
			break
		}
		pb.SampleIn = append(pb.SampleIn, string(file))

		file, err = ioutil.ReadFile(dir + strconv.Itoa(i) + ".out")
		if err != nil {
			log.Println(err)
			break
		}
		pb.SampleOut = append(pb.SampleOut, string(file))
	}
	c.JSON(http.StatusOK, pb)
}

// *******************************funcs
func initSample(dir string) {
	var err error
	for i := 1; ; i++ {
		err = os.Remove(dir + strconv.Itoa(i) + ".in")
		if err != nil {
			break
		}
		_ = os.Remove(dir + strconv.Itoa(i) + ".out")
	}
}

func isPrivate(probNo int) bool {
	var stat int
	Udb.QueryRow("select stat from probs where prob_no=?", probNo).Scan(&stat)
	return (stat == 0)
}

func moveFile(oriNo string) {
	err := os.Rename(privDir+oriNo, pubDir+oriNo)
	printErr(err)
}
