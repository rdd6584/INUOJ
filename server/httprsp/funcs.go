package httprsp

import (
	"crypto/rand"
	"math/big"
	"net/smtp"
)

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

func makeAuthKey() string {
	ret := ""
	for i := 0; i < 20; i++ {
		var x *big.Int
		x, _ = rand.Int(rand.Reader, big.NewInt(2))
		if x.Int64() == 1 {
			x, _ = rand.Int(rand.Reader, big.NewInt(26))
			ret += string(rune(x.Int64() + 65))
		} else {
			x, _ = rand.Int(rand.Reader, big.NewInt(100))
			ret += x.String()
		}
	}
	return ret
}

func sendMail(rcpt string) {
	auth := smtp.PlainAuth("", "inuojteam@gmail.com", "inu20190325", "smtp.gmail.com")

	from := "inuojteam@gmail.com"
	to := []string{rcpt}
	authkey := makeAuthKey()

	// 메시지 작성
	headerSubject := "Subject: INUOJ 이메일 주소 인증\r\n\r\n"
	body := "아래 링크를 누르시면 이메일 인증이 완료됩니다.\r\n"
	link := domain + "/auth?token=" + authkey + "&email=" + to[0]
	msg := []byte(headerSubject + body + link)

	// 메일 보내기
	var err error
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}

	var res bool
	err = Udb.QueryRow("SELECT NOT EXISTS (SELECT * FROM authtokens where email=?)", rcpt).Scan(&res)
	if err != nil {
		panic(err)
	}
	if res {
		_, err = Udb.Exec("insert into authtokes values(?, ?)", rcpt, authkey)
	} else {
		_, err = Udb.Exec("update authtokens set token=?, auth_time=current_timestamp where email=?", authkey, rcpt)
	}
	if err != nil {
		panic(err)
	}
}
