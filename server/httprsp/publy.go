package httprsp

const (
	mailSender string = "inuojteam@gmail.com"
	mailPass   string = "inu20190325"
	ourKey     string = "Rs9b6w22qGB1po1a7^$37agXez"
	OurMysql   string = "rdd:20190325@tcp(192.168.0.42:3306)/inuoj"
	privDir    string = "../Judger/problems/private/"
	pubDir     string = "../Judger/problems/public/"
	codeDir    string = "../Judger/usercodes/"
	sampleDir  string = "/data/sample"
	dataDir    string = "/data/"
	inDir      string = "/data/in/"
	outDir     string = "/data/out/"
	postDir    string = "../posts/"
)

const (
	Cpp = iota + 1
	C
	Go
	Java
	Pypy
	Python
	Kotlin
	LangSize
)

const (
	All = iota
	AC
	WA
	TLE
	MLE
	RE
	CE
	ServerError
	ResultSize
)

const (
	Notice = iota + 1
	Question
	Free
	CategorySize
)

const (
	Cdmin = iota
	Bdmin
	Admin
)

var descName = []string{"/main.txt", "/input.txt", "/output.txt"}
