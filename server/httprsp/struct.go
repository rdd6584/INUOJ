package httprsp

import "mime/multipart"

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

type editInfo struct {
	ID          string `json:"id" binding:"required"`
	Password    string `json:"pass" binding:"required"`
	NewPassword string `json:"newpass" binding:"required"`
}

type paramInfo struct {
	Name  string
	Type  int    // int: 0, string: 1
	Value string // int: "0", string: "" is null
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

//************************ start probs
type probDetail struct {
	OriNo       int      `json:"ori_no" form:"ori_no"`
	ProbNo      int      `json:"prob_no" form:"prob_no"` // 문제 등록시 부여
	TimeLimit   int      `json:"t_limit" form:"t_limit"`
	MemoryLimit int      `json:"m_limit" form:"m_limit"`
	Attempt     int      `json:"attempt" form:"attempt"`
	Accept      int      `json:"accept" form:"accept"`
	Owner       string   `json:"owner" form:"owner"`
	Title       string   `json:"title" form:"title"`
	Stat        int      `json:"stat" form:"stat"`               // 0 미등록, 1 채점 준비, 2 등록, 3 비공개
	Description []string `json:"description" form:"description"` // 본문, 입력, 출력
	SampleIn    []string `json:"samplein" form:"samplein"`
	SampleOut   []string `json:"sampleout" form:"sampleout"`
}

type probForList struct {
	OriNo   int    `json:"ori_no"`
	ProbNo  int    `json:"prob_no"`
	Title   string `json:"title"`
	Attempt int    `json:"attempt"`
	Accept  int    `json:"accept"`
	Stat    int    `json:"stat"`
}

type probData struct {
	OriNo  int                     `form:"ori_no" binding:"required"`
	Input  []*multipart.FileHeader `form:"input"`
	Output []*multipart.FileHeader `form:"output"`
}

type fileArray struct {
	OriNo    int      `json:"ori_no" binding:"required"`
	FileList []string `json:"files"`
}
