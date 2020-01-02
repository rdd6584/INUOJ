package httprsp

import (
	"github.com/gin-gonic/gin"
)

func ResRouter(e *gin.Engine) {
	app := e.Group("/")
	{
		app.GET("/", toMain)
		app.GET("/login", login)
		app.GET("/register", register)
		app.POST("/regi-done", regiComplete)  // id, pass, email 정보 저장
		app.GET("/register/valid", regiValid) // id or email의 unique 여부, status : ok
		app.POST("/login/valid", loginSuc)    // id, pass 일치 여부, status : ok
	}
}
