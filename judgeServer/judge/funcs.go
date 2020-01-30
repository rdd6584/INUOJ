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

func updateResultList(submNo int, ac bool) {
	var probNo int
	var userID string

	// 트랜젝션 해야할까 ?
	err := Udb.QueryRow("select prob_no, id from submits where subm_no=?", submNo).Scan(&probNo, &userID)
	printErr(err)
	var ex int
	err = Udb.QueryRow("select result from result_list where id=? and prob_no=?", userID, probNo).Scan(&ex)
	if err != nil {
		_, err = Udb.Exec("insert into result_list values(?,?,?)", userID, probNo, WA)
		printErr(err)
		_, err = Udb.Exec("update user_info set wa_count=wa_count+1 where id=?", userID)
		printErr(err)
		_, err = Udb.Exec("update probs set attempt=attempt+1 where prob_no=?", probNo)
		printErr(err)
	} else if ac && ex == WA {
		_, err = Udb.Exec("update result_list set result=? where id=? and prob_no=?", AC, userID, probNo)
		printErr(err)
		_, err = Udb.Exec("update user_info set wa_count=wa_count-1, ac_count=ac_count+1 where id=?", userID)
		printErr(err)
		_, err = Udb.Exec("update probs set accept=accept+1 where prob_no=?", probNo)
		printErr(err)
	}

	if !ac {
		_, err = Udb.Exec("update user_info set all_count=all_count+1 where id=?", userID)
		printErr(err)
	}
}

func ReadQueue() {
	result, _ := Udb.Query("select * from judge_q")
	defer result.Close()
	var submNo, oriNo, lang int
	var err error

	for result.Next() {
		err = result.Scan(&submNo, &oriNo, &lang)
		if err != nil {
			log.Println("queue : ", err)
			continue
		}
		updateResultList(submNo, false)

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
	inputDir := pubDir + strconv.Itoa(oriNo) + inDir
	dataDir := pubDir + strconv.Itoa(oriNo) + outDir

	files, err := ioutil.ReadDir("../Judger/problems/public/" + strconv.Itoa(oriNo) + inDir)
	printErr(err)

	var script string
	var maxTime, maxMem int
	for _, file := range files {
		script = "./libjudger.so " + "--max_cpu_time=" + strconv.Itoa(proT) +
			" --max_real_time=" + strconv.Itoa(proT) + " --max_memory=" + strconv.Itoa(proM*1024*1024) +
			" --max_process_number=" + maxProcessNum + " --max_output_size=" + maxOutputSize +
			" --exe_path=" + getExeFile(lang) + getArgs(lang) + " --input_path=" + inputDir + file.Name() +
			" --output_path=./output.txt" + " --error_path=./error.txt" +
			" --uid=0 --gid=0" + seccompRule(lang)

		c := exec.Command("/bin/bash", "-c", script)
		stdout, err := c.CombinedOutput()
		printErr(err)

		var data judgeResult
		json.Unmarshal(stdout, &data)

		if data.Result != 0 {
			if data.Result == 4 {
				if data.Signal == 25 {
					log.Println("pass")
					setStatus(submNo, WA) // 출력 초과
				} else {
					setStatus(submNo, RE) // 런타임에러
				}
			} else if data.Result == 1 || data.Result == 2 {
				setStatus(submNo, TLE) // 시간초과
			} else if data.Result == 3 {
				setStatus(submNo, MLE) // 메모리초과
			} else if data.Result == 5 {
				setStatus(submNo, ServerError) // 서버 에러
			}
			return
		}

		outfile := toOut(file.Name())
		userOutput, _ := ioutil.ReadFile("./output.txt")
		solOutput, _ := ioutil.ReadFile(dataDir + outfile)

		userOutputArr := strings.Split(strings.TrimRight(string(userOutput), " \n"), "\n")
		solOutputArr := strings.Split(strings.TrimRight(string(solOutput), " \n"), "\n")

		if len(userOutputArr) != len(solOutputArr) {
			setStatus(submNo, WA)
			return
		}
		for i := len(userOutputArr) - 1; i >= 0; i-- {
			if strings.TrimRight(userOutputArr[i], " ") != strings.TrimRight(solOutputArr[i], " ") {
				setStatus(submNo, WA)
				return
			}
		}
		maxTime = getMax(data.RealTime, maxTime)
		maxTime = getMax(data.CPUTime, maxTime)
		maxMem = getMax(data.Memory, maxMem)
	}
	setStatus(submNo, AC, maxTime, maxMem/1024)
	updateResultList(submNo, true)
}

func compile(lang int, submNo string) bool {
	copyFile(submitDir+submNo+"/Main"+fileType(lang), "./Main"+fileType(lang))

	var script string
	switch lang {
	case C:
		script = "gcc Main.c -o Main.o -O2 -Wall -lm -static -std=c99 -DONLINE_JUDGE -DBOJ"
	case Cpp:
		script = "g++ Main.cpp -o Main.o -O2 -Wall -lm -static -std=gnu++17 -DONLINE_JUDGE -DBOJ"
	case Java:
		script = "javac -J-Xms1024m -J-Xmx1024m -J-Xss512m -encoding UTF-8 Main.java"
	case Python, Pypy:
		script = "python3 -m py_compile Main.py"
	}

	c := exec.Command("/bin/bash", "-c", script)
	stdout, err := c.CombinedOutput()
	ioutil.WriteFile(submitDir+submNo+"/compileMsg.txt", stdout, 0644)
	if err != nil {
		return false
	}
	return true
}

func seccompRule(lang int) string {
	switch lang {
	case C, Cpp:
		return " --seccomp_rule_name=c_cpp"
	case Python:
		return " --seccomp_rule_name=general"
	}
	return ""
}

func getExeFile(lang int) string {
	var ret string
	switch lang {
	case C, Cpp:
		ret = "Main.o"
	case Java:
		ret = "/usr/bin/java"
	case Python:
		ret = "/usr/bin/python3"
	case Pypy:
		ret = "/usr/bin/pypy3"
	}
	return ret
}

func getArgs(lang int) string {
	switch lang {
	case Java:
		return " --args=-Dfile.encoding=UTF-8 --args=Main --memory_limit_check_only=1"
	case Python, Pypy:
		return " --args=Main.py"
	}
	return ""
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
	case Java:
		return ".java"
	case Python, Pypy:
		return ".py"
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

func copyFile(src, dst string) {
	content, err := ioutil.ReadFile(src)
	printErr(err)
	err = ioutil.WriteFile(dst, content, 0644)
	printErr(err)
}
