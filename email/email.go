package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

//SMTP smtp auth info
type SMTP struct {
	address,
	username,
	password,
	authtype string
}

//New return *SMTP
func New(address, username, password, authtype string) *SMTP {
	return &SMTP{
		address:  address,
		username: username,
		password: password,
		authtype: authtype,
	}
}

//SendMail sent the email with args
func (s *SMTP) SendMail(from string, to []string, subject, body, contentType string) error {

	tos := strings.Join(to, ";")
	addrArr := strings.Split(s.address, ":")
	if len(addrArr) != 2 {
		return fmt.Errorf("address format error")
	}

	//base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	b64 := base64.StdEncoding
	header := make(map[string]string)

	header["From"] = s.username
	if strings.HasSuffix(from, "staff.qkagame.com") {
		header["From"] = from
	}
	header["To"] = tos
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))
	header["MIME-Version"] = "1.0"

	if contentType == "html" {
		header["Content-Type"] = "text/html; charset=UTF-8"
	} else {
		header["Content-Type"] = "text/plain; charset=UTF-8"
	}

	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))
	var auth smtp.Auth
	switch s.authtype {
	case "LOGIN":
		auth = LoginAuth(s.username, s.password)
	case "CRAM-MD5":
		auth = smtp.CRAMMD5Auth(s.username, s.password)
	default:
		auth = smtp.PlainAuth("", s.username, s.password, addrArr[0])

	}

	return smtp.SendMail(s.address, auth, s.username, to, []byte(message))
}
