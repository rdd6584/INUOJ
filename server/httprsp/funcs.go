package httprsp

import "net/smtp"

func isNotNull(x paramInfo) bool {
	if x.Type == 0 {
		if x.Value == "-1" {
			return false
		}
	} else if x.Type == 1 {
		if x.Value == "" {
			return false
		}
	}
	return true
}

func addParam(str *string, x paramInfo) {
	if x.Type == 0 {
		*str += x.Name + "=" + x.Value + " "
	} else if x.Type == 1 {
		*str += x.Name + "='" + x.Value + "' "
	}
}

func makeWhere(x ...paramInfo) string {
	ret := ""
	fst := true
	for _, cur := range x {
		if isNotNull(cur) {
			if fst {
				ret += "where "
				fst = false
			} else {
				ret += "and "
			}
			addParam(&ret, cur)
		}
	}
	return ret
}

func sendMail() {
	auth := smtp.PlainAuth("", "@gmail.com", "pwd", "smtp.gmail.com")

	from := "orihehe@gmail.com"
	to := []string{"rdd573@naver.com"} // 복수 수신자 가능

	// 메시지 작성
	headerSubject := "Subject: 테스트\r\n"
	headerBlank := "\r\n"
	body := "메일 테스트입니다\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
