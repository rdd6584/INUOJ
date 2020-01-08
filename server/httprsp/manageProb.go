package httprsp

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var fileName = []string{"/main.txt", "/input.txt", "/output.txt"}

func getNewOriNo(c *gin.Context) {
	var ret int
	err := Udb.QueryRow("select max(ori_no)").Scan(&ret)
	printErr(err)
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

	dir := probDir + strconv.Itoa(pb.ProbNo)
	_ = os.Mkdir(dir, os.ModePerm)
	_ = os.Mkdir(dir+dataDir, os.ModePerm)
	for i := 0; i < len(fileName); i++ {
		err = ioutil.WriteFile(dir+fileName[i], []byte(pb.Description[i]), 0644)
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
	if len(pb.Input) != len(pb.Output) {
		c.String(http.StatusBadRequest, "in, out not match")
		return
	}

	dir := probDir + strconv.Itoa(pb.OriNo) + dataDir + "/"
	_ = os.Mkdir(dir, os.ModePerm)
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
	dir := probDir + strconv.Itoa(files.OriNo) + dataDir + "/"
	for _, file := range files.FileList {
		_ = os.Remove(dir + file + ".in")
		_ = os.Remove(dir + file + ".out")
	}
	c.String(http.StatusOK, "")
}

func viewProbDetail(c *gin.Context) {
	var pNum probView
	var pb probDetail
	var err error
	if err = c.ShouldBind(&pNum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(pNum.ProbNum)
	err = Udb.QueryRow("select * from probs where prob_no=?", pNum.ProbNum).Scan(&pb.OriNo, &pb.ProbNo,
		&pb.TimeLimit, &pb.MemoryLimit, &pb.Attempt, &pb.Accept, &pb.Owner, &pb.Title, &pb.Stat)
	if err != nil {
		log.Println(err)
	}

	dir := probDir + strconv.Itoa(pb.ProbNo)
	var file []byte
	for i := 0; i < len(fileName); i++ {
		file, err = ioutil.ReadFile(dir + fileName[i])
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
	c.SecureJSON(http.StatusOK, pb)
}

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
