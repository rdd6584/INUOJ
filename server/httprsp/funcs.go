package httprsp

import (
	"crypto/rand"
	"database/sql"
	"log"
	"math/big"
	"net/smtp"
)

func affectedOneRow(result sql.Result) bool {
	n, err := result.RowsAffected()
	printErr(err)
	return (n == 1)
}

func isAuthUser(userID string) bool {
	var auth bool
	err := Udb.QueryRow("select auth from users where id=?", userID).Scan(&auth)
	if err != nil || !auth {
		return false
	}
	return true
}

func fileType(lang int) string {
	switch lang {
	case Cpp:
		return ".cpp"
	case C:
		return ".c"
	case Java:
		return ".java"
	}
	return ""
}

func whatDmin(userID string, dmin int) bool {
	var res int
	err := Udb.QueryRow("select admin from users where id=?", userID).Scan(&res)
	printErr(err)
	return (res >= dmin)
}

func isCorrectInfo(userID string, userPass string) bool {
	var dbPass string
	err := Udb.QueryRow("select password from users where id=?", userID).Scan(&dbPass)
	printErr(err)
	return (dbPass == userPass)
}

func userAuthValid(userID string) (bool, int) {
	var auth bool
	var admin int
	err := Udb.QueryRow("select auth, admin from users where id=?", userID).Scan(&auth, &admin)
	printErr(err)
	return auth, admin
}

func editPass(userID string, newPass string) {
	_, err := Udb.Exec("update users set password=? where id=?", newPass, userID)
	printErr(err)
}

func isNotNull(x paramInfo) bool {
	if x.Type == 0 {
		if x.Value == "0" {
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
	auth := smtp.PlainAuth("", mailSender, mailPass, "smtp.gmail.com")

	from := mailSender
	to := []string{rcpt}
	authkey := makeAuthKey()

	// 메시지 작성
	headerSubject := "To: " + rcpt + "\r\nSubject: INUOJ 이메일 주소 인증\r\n\r\n"
	body := "아래 링크를 누르시면 이메일 인증이 완료됩니다.\r\n"
	link := domain + "/auth?token=" + authkey + "&email=" + to[0]
	msg := []byte(headerSubject + body + link)

	var cnt int
	var err error

	tx, err := Udb.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_ = tx.QueryRow("select count from authtokens where email=?", rcpt).Scan(&cnt)
	if cnt == 0 {
		_, err = tx.Exec("insert into authtokens(email, token) values(?, ?)", rcpt, authkey)
	} else {
		_, err = tx.Exec("update authtokens set token=?, auth_time=current_timestamp, count=? where email=?", authkey, cnt, rcpt)
	}
	if err != nil {
		log.Panic(err)
	}

	// 메일 보내기
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func printErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
func panicErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
