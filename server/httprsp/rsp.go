package httprsp

import (
	"github.com/gin-gonic/gin"
)

func ResRouter(e *gin.Engine) {
	var authAll = initJWT(loginUserAuthorizator)

	app := e.Group("/api")
	{
		app.POST("/regi-done", regiComplete)  // id, pass, email 정보 저장
		app.GET("/register/valid", regiValid) // id or email의 unique 여부, status : 1
		app.POST("/login/valid", authAll.LoginHandler)
		app.POST("/mailauth", mailAuth)
		app.POST("/sendauth", reSendMail) // id받고 status=true : 이미 등록됨, false : 메일 전송함
		app.GET("/refresh", authAll.RefreshHandler)
	}
	app1 := e.Group("/api")
	app1.Use(authAll.MiddlewareFunc())
	{
		app1.GET("/status", getStatus) // data_num : 전체 데이터 개수, datas : 제출기록
		app1.GET("/logout", authAll.LogoutHandler)
	}

	e.NoRoute(toMain)
}
