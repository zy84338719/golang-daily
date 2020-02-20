package main

import (
	"fmt"
	"mail/utils"
)

func main() {
	tencent := utils.ParseFile("mail/config/config.json")
	fmt.Println("send email")
	subject, body := createMail()
	to := "zy84338719@hotmail.com"
	err := utils.SendToMail(tencent, to, subject, body, "text/html")
	if err != nil {
		fmt.Println("Send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("Send mail success!")
	}
}

func createMail() (subject string, body string) {
	subject = "使用Golang发送邮件"

	body = `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
</>
		</body>
		</html>
		`
	return
}
