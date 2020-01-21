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
	Udb.QueryRow("select stat from probs where ori_no=?", json.OriNo).Scan(&stat)
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
		moveFile(strconv.Itoa(json.OriNo))
		// 트랜젝션 ?
		var probNo int
		err = Udb.QueryRow("select max(prob_no) from probs").Scan(&probNo)
		printErr(err)

		_, err = Udb.Exec("update probs set prob_no=? where ori_no=?", probNo+1, json.OriNo)
		printErr(err)
	}
	c.String(http.StatusOK, "")
}

func myProbList(c *gin.Context) {
	id := c.Query("id")
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

	var dir string
	if pb.Stat == 0 {
		dir = privDir + strconv.Itoa(pb.OriNo)
	} else {
		dir = pubDir + strconv.Itoa(pb.OriNo)
	}
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

	dir := privDir + strconv.Itoa(pb.OriNo)
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

	dir := privDir + strconv.Itoa(files.OriNo)
	for _, file := range files.FileList {
		_ = os.Remove(dir + inDir + file + ".in")
		_ = os.Remove(dir + outDir + file + ".out")
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

		// stat에 따라
		files, err := ioutil.ReadDir(pubDir + ori + inDir)
		printErr(err)

		for _, file := range files {
			fileName := file.Name()
			if !strings.HasPrefix(fileName, "sample") {
				pb.Datas = append(pb.Datas, fileName[0:len(fileName)-3])
			}
		}
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

	dir := pubDir + strconv.Itoa(pNum)
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
