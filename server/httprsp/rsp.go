package httprsp

import (
	"github.com/gin-gonic/gin"
)

func ResRouter(e *gin.Engine) {
	var authAll = initJWT(loginUserAuthorizator)

	e.GET("/ttt", myProbList)
	// *************** for all ***************
	app := e.Group("/api")
	{
		app.POST("/regi-done", regiComplete)     // id, pass, email 정보 저장
		app.GET("/register/valid", regiValid)    // id or email의 unique 여부
		app.POST("/login", authAll.LoginHandler) // 로그인기능, 인증 받은 계정만
		app.POST("/emailauth", emailAuth)        // 이메일 인증 링크 눌렀을 때의 과정
		app.POST("/sendauth", reSendMail)        // 인증 메일 다시 보내기, 최대 5회
		app.GET("/refresh", authAll.RefreshHandler)
	}

	// *************** auth 1 ***************
	app1 := e.Group("/api")
	app1.Use(authAll.MiddlewareFunc())
	{
		app1.GET("/status", getStatus)                       // 전체 데이터 개수, 제출기록
		app1.GET("/logout", authAll.LogoutHandler)           //***한번 보기
		app1.GET("/problem/detail/:prob_no", viewProbDetail) // 문제 디테일
		app1.POST("/problem/submit", probSubmit)
	}

	// *************** auth 1 && only me ***************
	var authMy = initJWT(onlyMeAuthorizator)
	app2 := e.Group("/api")
	app2.Use(authMy.MiddlewareFunc())
	{
		app2.POST("/edit", editUserInfo)
	}

	// *************** admin 1 ***************
	var authBdmin = initJWT(bdminAuthorizator)
	app3 := e.Group("/api/bdmin")
	app3.Use(authBdmin.MiddlewareFunc())
	{
		app3.GET("/new", getNewOriNo)               // 문제 추가
		app3.POST("/upload/desc", uploadDesc)       // 문제 디스크립션
		app3.POST("/upload/data", uploadData)       // 문제 데이터 추가
		app3.POST("/discard/data", discardData)     // 문제 데이터 삭제
		app3.GET("/detail/:ori_no", viewProbDetail) // 문제 디테일
		app3.GET("/list", myProbList)               // 권한이 있는 문제 목록
		app3.POST("/update/stat", changeStat)       // 문제 공개 상태 변경
	}

	e.NoRoute(toMain)
}
