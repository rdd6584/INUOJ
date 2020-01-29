package judge

const (
	OurMysql      string = "rdd:20190325@tcp(192.168.0.42:3306)/inuoj"
	dockerDir     string = "/home"
	hostDir       string = "/c/Users/rdd65/desktop/GoApp/src/INUOJ/Judger"
	privDir       string = "/problems/private/"
	pubDir        string = "/problems/public/"
	submitDir     string = "/usercodes/"
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
