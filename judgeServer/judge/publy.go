package judge

const (
	OurMysql      string = "root:ojbydduck0325!@tcp(192.168.0.42:3306)/inuoj"
	judgerDir     string = "../Judger"
	privDir       string = "../Judger/problems/private/"
	pubDir        string = "../Judger/problems/public/"
	submitDir     string = "../Judger/usercodes/"
	inDir         string = "/data/in/"
	outDir        string = "/data/out/"
	maxProcessNum string = "100"
	maxOutputSize string = "16384"
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
