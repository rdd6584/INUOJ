package judge

const (
	OurMysql      string = "root:ojbydduck0325!@tcp(127.0.0.1:3306)/inuoj"
	privDir       string = "problems/private/"
	pubDir        string = "problems/public/"
	submitDir     string = "usercodes/"
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
