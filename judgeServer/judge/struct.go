package judge

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

type judgeResult struct {
	CPUTime  int `json:"cpu_time"`
	RealTime int `json:"real_time"`
	Error    int `json:"error"`
	ExitCode int `json:"exit_code"`
	Memory   int `json:"memory"`
	Result   int `json:"result"`
	Signal   int `json:"signal"`
}
