package httprsp

import "mime/multipart"

const (
	pageSize int    = 15
	domain   string = "localhost"
)

type user struct {
	ID    string `json:"id"`
	Admin int    `json:"admin"`
}

type regiInfo struct {
	ID       string `form:"id" json:"id"`
	Password string `form:"pass" json:"pass"`
	Email    string `form:"email" json:"email"`
}

type loginInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
}

type loginRes struct {
	Status bool `json:"status"`
	Auth   bool `json:"auth"`
}

type editPassword struct {
	ID          string `json:"id" binding:"required"`
	Password    string `json:"pass" binding:"required"`
	NewPassword string `json:"newpass" binding:"required"`
}

type editPR struct {
	ID string `json:"id" binding:"required"`
	PR string `json:"pr"`
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
	SubmNo    int    `json:"subm_no"`
	ID        string `json:"id"`
	ProbNo    int    `json:"prob_no"`
	ProbTitle string `json:"prob_title"`
	Result    int    `json:"result"`
	Lang      int    `json:"lang"`
	RunTime   int    `json:"run_time"`
	Memory    int    `json:"memory"`
	Codelen   int    `json:"codelen"`
	SubmTime  string `json:"subm_time"`
}

type newSubmit struct {
	ID       string `json:"id" binding:"required"`
	ProbNo   int    `json:"prob_no" binding:"required"`
	Lang     int    `json:"lang" binding:"required"`
	UserCode string `json:"code" binding:"required"`
}

type submitPage struct { // status page
	DataNum int          `json:"data_num"`
	Datas   []submitInfo `json:"datas"`
}

type probListPage struct { // prob list page
	DataNum  int           `json:"data_num"`
	Problems []probListAll `json:"problems"`
}

type userPage struct {
	ID       string `json:"id"`
	PR       string `json:"pr"`
	ACcount  int    `json:"ac_count"`
	WAcount  int    `json:"wa_count"`
	ALLcount int    `json:"all_count"`
	Rank     int    `json:"rank"`
}

type rankPage struct {
	ID      string `json:"id"`
	PR      string `json:"pr"`
	ACcount int    `json:"ac_count"`
	Rank    int    `json:"rank"`
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
	Datas       []string `json:"datas"`
}

type probListMy struct {
	OriNo  int    `json:"ori_no"`
	ProbNo int    `json:"prob_no"`
	Title  string `json:"title"`
	Stat   int    `json:"stat"`
}

type probListAll struct {
	ProbNo  int    `json:"prob_no"`
	Title   string `json:"title"`
	Attempt int    `json:"attempt"`
	Accept  int    `json:"accept"`
	Stat    int    `json:"stat"`
	Result  int    `json:"result"`
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

type modStat struct {
	OriNo    int `json:"ori_no"`
	FromStat int `json:"fromstat"`
	ToStat   int `json:"tostat"`
}

//************************ start board
type postInfo struct {
	PostNo    int       `json:"post_no"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	Code      string    `json:"code"`
	ID        string    `json:"id" binding:"required"`
	Category  int       `json:"category"`
	ProbNo    int       `json:"prob_no"`
	ProbTitle string    `json:"prob_title"`
	CmtNo     int       `json:"cmt_no"`
	CmtList   []cmtInfo `json:"cmt_list"`
	Result    int       `json:"result"`
	PostTime  string    `json:"post_time"`
}

type cmtInfo struct {
	PostNo  int    `json:"post_no" binding:"required"`
	ID      string `json:"id" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	CmtTime string `json:"cmt_time"`
}

type noticeInfo struct {
	PostNo int  `json:"post_no" binding:"required"`
	Notice bool `json:"notice"`
}
