package judge

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

var Udb *sql.DB
var Ch chan struct{}

func ReadQueue() {
	result, err := Udb.Query("select * from judge_q")
	var submNo, oriNo, lang int
	for result.Next() {
		err = result.Scan(&submNo, &oriNo, &lang)
		if err != nil {
			log.Println("queue : ", err)
			break
		}
		if !compile(lang, strconv.Itoa(submNo)) {
			// 컴파일 에러
			continue
		}
		run(oriNo, lang)
	}
}

func run(oriNo int, lang int) {
	var proT, proM, stat int
	err := Udb.QueryRow("select t_limit, m_limit, stat from probs where ori_no=?",
		oriNo).Scan(&proT, &proM, &stat)
	printErr(err)

	// stat에 따라 경로 설정
	inputDir := dockerDir + pubDir + strconv.Itoa(oriNo) + inDir

	files, err := ioutil.ReadDir("../Judger/problems/public/" + strconv.Itoa(oriNo) + inDir)
	printErr(err)

	for _, file := range files {
		script := "docker run --rm -v " + hostDir + ":" + dockerDir + " " +
			"gcc " + dockerDir + "/libjudger.so --max_cpu_time=" + strconv.Itoa(proT) +
			" --max_real_time=" + strconv.Itoa(proT) + " --max_memory=" + strconv.Itoa(proM*1024*1024) +
			" --max_process_number=" + maxProcessNum + " --max_output_size=" + maxOutputSize +
			" --exe_path=\"" + dockerDir + "/Main.o" + "\" --input_path=\"" + inputDir + file.Name() +
			"\" --output_path=\"" + dockerDir + "/output.txt" + "\" --error_path=\"" + dockerDir + "/error.txt" +
			"\" --uid=0 --gid=0 --seccomp_rule_name=\"c_cpp\""

		args := script[0:len(script)]
		c := exec.Command("cmd", "/C", args)
		log.Println(c)
		stdout, err := c.CombinedOutput()
		printErr(err)
		log.Println(string(stdout))
	}

}

func compile(lang int, submNo string) bool {
	script := "docker run --rm -v " + hostDir + ":" + dockerDir + " "
	switch lang {
	case C:
		script += "gcc gcc /home" + submitDir + submNo +
			".c -o /home/Main.o -O2 -Wall -lm -static -std=c99 -DONLINE_JUDGE -DBOJ"
	case Cpp:
		script += "gcc g++ /home" + submitDir + submNo +
			".cpp -o /home/Main.o -O2 -Wall -lm -static -std=gnu++17 -DONLINE_JUDGE -DBOJ"
	}
	args := script[0:len(script)]
	c := exec.Command("cmd", "/C", args)
	stdout, err := c.CombinedOutput()
	ioutil.WriteFile(submNo+".txt", stdout, 0644)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func fileType(lang int) string {
	switch lang {
	case Cpp:
		return ".cpp"
	case C:
		return ".c"
	}
	return ""
}

func printErr(e error) {
	if e != nil {
		log.Println(e)
	}
}

func panicErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
