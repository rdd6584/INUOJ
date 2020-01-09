package httprsp

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getNewOriNo(c *gin.Context) {
	var ret int
	var err error
	id := c.Query("id")
	tx, err := Udb.Begin()
	panicErr(err)
	defer func() {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "")
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "in, out not match"})
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
	var pNum int
	var pb probDetail
	var err error
	pNum, err = strconv.Atoi(c.Param("ori_no"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = Udb.QueryRow("select t_limit, m_limit, owner, title from probs where ori_no=?", pNum).Scan(
		&pb.TimeLimit, &pb.MemoryLimit, &pb.Owner, &pb.Title)
	if err != nil {
		log.Println(err)
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
