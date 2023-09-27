package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "PeyChou <peychou@qq.com>"
	e.To = []string{"zpx2503540980@gmail.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为<h1>123456</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "peychou@qq.com", "foyglitzssradjbd", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
