package httprsp

const (
	pageSize int    = 15
	domain   string = "localhost"
)

type user struct {
	ID string `json:"id"`
}

type regiInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type loginInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
}

type loginRes struct {
	Status bool `json:"status"`
	Auth   bool `json:"auth"`
}

type paramInfo struct {
	Name  string
	Type  int    // int: 0, string: 1
	Value string // int: -1, string: "" is null
}

type authInfo struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type submitInfo struct {
	SubmNo   int    `json:"subm_no"`
	ID       string `json:"id"`
	ProbNo   int    `json:"prob_no"`
	Result   int    `json:"result"`
	Lang     int    `json:"lang"`
	RunTime  int    `json:"run_time"`
	Memory   int    `json:"memory"`
	Codelen  int    `json:"codelen"`
	SubmTime string `json:"subm_time"`
}

type submitPage struct {
	DataNum int          `json:"data_num"`
	Datas   []submitInfo `json:"datas"`
}
