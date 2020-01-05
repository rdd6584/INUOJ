package httprsp

import (
	"github.com/gin-gonic/gin"
)

func ResRouter(e *gin.Engine) {
	e.GET("/", toMain)
	e.POST("/regi-done", regiComplete)  // id, pass, email 정보 저장
	e.GET("/register/valid", regiValid) // id or email의 unique 여부, status : 1
	e.POST("/login/valid", loginSuc)    // id, pass 일치 여부, status : true, auth : true
	e.GET("/status", getStatus)         // data_num : 전체 데이터 개수, datas : 제출기록
	e.POST("/auth-done", authComplete)

}
