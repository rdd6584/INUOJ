package httprsp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func getUserInfo(c *gin.Context) {
	userID := c.Param("userid")
	if !isAuthUser(userID) {
		c.String(http.StatusNotFound, "")
		return
	}
	var json userPage
	err := Udb.QueryRow("select * from (select *, rank() over (order by ac_count desc) "+
		"as ranking from user_info) as rr where rr.id=?", userID).Scan(&json.ID,
		&json.PR, &json.AC_Count, &json.WA_Count, &json.ALL_Count, &json.Rank)
	printErr(err)

	c.JSON(http.StatusOK, json)
}

func getUserProbList(c *gin.Context) {
	userID := c.Query("id")
	result := c.Query("result")
	if userID == "" || result == "" {
		c.String(http.StatusBadRequest, "")
		return
	}
	// result string 가능한지 보자
	rows, err := Udb.Query("select pr.title, pr.prob_no from probs as pr join "+
		"result_list as rl where pr.prob_no=rl.prob_no and rl.id=? and rl.result=?", userID, result)
	printErr(err)

	var nums []int
	var titles []string
	var num int
  var title string
	for rows.Next() {
		err = rows.Scan(&title, &num)
		printErr(err)
		nums = append(nums, num)
		titles = append(titles, num)
	}
	c.JSON(http.StatusOK, gin.H{"numbers": nums, "titles": titles})
}

func editUserInfo(c *gin.Context) {
	var json editInfo
	var err error
	if err = c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !isCorrectInfo(json.ID, json.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail"})
		return
	}
	editPass(json.ID, json.NewPassword)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
