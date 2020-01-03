package httprsp

type regiInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}
type loginInfo struct {
	ID       string `form:"id" json:"id" binding:"required"`
	Password string `form:"pass" json:"pass" binding:"required"`
}
type paramInfo struct {
	Name  string
	Type  int    // int: 0, string: 1
	Value string // int: -1, string: "" is null
}
