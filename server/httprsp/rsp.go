package httprsp

import (
	"github.com/gin-gonic/gin"
)

func ResRouter(e *gin.Engine) {
	var authMiddleware = initJWT()

	app := e.Group("/api")
	{
		app.POST("/regi-done", regiComplete)                  // id, pass, email 정보 저장
		app.GET("/register/valid", regiValid)                 // id or email의 unique 여부, status : 1
		app.POST("/login/valid", authMiddleware.LoginHandler) // id, pass 일치 여부, status : true, auth : true, token="a1b2"
		app.POST("/mailauth", mailAuth)
		app.GET("/refresh", authMiddleware.RefreshHandler)
	}
	app1 := e.Group("/api")
	app1.Use(authMiddleware.MiddlewareFunc())
	{
		app1.GET("/status", getStatus) // data_num : 전체 데이터 개수, datas : 제출기록
		app1.GET("/logout", authMiddleware.LogoutHandler)
	}
	e.NoRoute(toMain)
}
