package judge

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

var Udb *sql.DB

func ReadQueue() {
	result, err := Udb.Query("select * from judge_q")
	var submNo, oriNo, lang int
	for result.Next() {
		err = result.Scan(&submNo, &oriNo, &lang)
		if err != nil {
			log.Println("queue : ", err)
			break
		}
		_, _ = Udb.Exec("delete from judge_q where subm_no=?", submNo)
		if !compile(lang, strconv.Itoa(submNo)) {
			setStatus(submNo, CE) // 컴파일 에러
			continue
		}
		run(oriNo, lang, submNo)
	}
}

func setStatus(submNo int, args ...int) {
	if args[0] == 1 {
		_, _ = Udb.Exec("update submits set result=?, run_time=?, memory=? where subm_no=?",
			args[0], args[1], args[2], submNo)
	} else {
		_, _ = Udb.Exec("update submits set result=? where subm_no=?", args[0], submNo)
	}
}

func run(oriNo int, lang int, submNo int) {
	var proT, proM, stat int
	err := Udb.QueryRow("select t_limit, m_limit, stat from probs where ori_no=?",
		oriNo).Scan(&proT, &proM, &stat)
	printErr(err)

	// stat에 따라 경로 설정해야함
	inputDir := dockerDir + pubDir + strconv.Itoa(oriNo) + inDir
	dataDir := "../Judger" + pubDir + strconv.Itoa(oriNo) + outDir

	files, err := ioutil.ReadDir("../Judger/problems/public/" + strconv.Itoa(oriNo) + inDir)
	printErr(err)

	var script string
	var maxTime, maxMem int
	for _, file := range files {
		script = "docker run --rm -v " + hostDir + ":" + dockerDir + " " +
			"gcc " + dockerDir + "/libjudger.so --max_cpu_time=" + strconv.Itoa(proT) +
			" --max_real_time=" + strconv.Itoa(proT) + " --max_memory=" + strconv.Itoa(proM*1024*1024) +
			" --max_process_number=" + maxProcessNum + " --max_output_size=" + maxOutputSize +
			" --exe_path=" + dockerDir + "/Main.o" + " --input_path=" + inputDir + file.Name() +
			" --output_path=" + dockerDir + "/output.txt" + " --error_path=" + dockerDir + "/error.txt" +
			" --uid=0 --gid=0 --seccomp_rule_name=c_cpp"

		c := exec.Command("cmd", "/C", script)
		stdout, err := c.CombinedOutput()
		printErr(err)

		var data judgeResult
		json.Unmarshal(stdout, &data)

		//log.Println(data)

		if data.Result != 0 {
			if data.Result == 4 {
				if data.Signal == 25 {
					setStatus(submNo, WA) // 출력 초과
				} else {
					setStatus(submNo, RE) // 런타임에러
				}
			} else if data.Result == 1 || data.Result == 2 {
				setStatus(submNo, TLE) // 시간초과
			} else if data.Result == 3 {
				setStatus(submNo, MLE) // 메모리초과
			} else if data.Result == 5 {
				setStatus(submNo, SERVER_ERR) // 서버 에러
			}
			return
		}

		outfile := toOut(file.Name())
		userOutput, _ := ioutil.ReadFile("../Judger/output.txt")
		solOutput, _ := ioutil.ReadFile(dataDir + outfile)

		userOutputArr := strings.Split(strings.TrimRight(string(userOutput), " \n"), "\n")
		solOutputArr := strings.Split(strings.TrimRight(string(solOutput), " \n"), "\n")

		//log.Println(file.Name(), userOutputArr[0], solOutputArr[0])
		if len(userOutputArr) != len(solOutputArr) {
			setStatus(submNo, WA)
			return
		}
		for i := len(userOutputArr) - 1; i >= 0; i-- {
			if userOutputArr[i] != solOutputArr[i] {
				setStatus(submNo, WA)
				return
			}
		}
		maxTime = getMax(data.RealTime, maxTime)
		maxTime = getMax(data.CPUTime, maxTime)
		maxMem = getMax(data.Memory, maxMem)
	}
	setStatus(submNo, AC, maxTime, maxMem)
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
	c := exec.Command("cmd", "/C", script)
	stdout, err := c.CombinedOutput()
	ioutil.WriteFile("../Judger/usercodes/"+submNo+".txt", stdout, 0644)
	if err != nil {
		log.Println("compile err : ", err)
		return false
	}
	return true
}

func toOut(name string) string {
	return name[0:len(name)-2] + "out"
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

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
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
