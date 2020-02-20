package utils

import (
	"crypto/tls"
	"encoding/json"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"mail/model"
)

func SendToMail(tencent model.Tencent, to, subject, body, mailtype string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", tencent.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody(mailtype, body)
	//m.Attach("/Users/zc/Pictures/skip.gif")
	d := gomail.NewDialer(tencent.Host, tencent.Port, tencent.User, tencent.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return err
	}
	return nil
}

func ParseFile(filename string) (mailConfig model.Tencent) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &mailConfig)
	if err != nil {
		return
	}
	log.Printf("反序列化成功  conf: %#v\n  host: %#v\n", mailConfig, mailConfig.Host)
	return
}
