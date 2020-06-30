package email

//github.com/go-gomail/gomail/blob/master/auth.go
import (
	"bytes"
	"fmt"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

//LoginAuth return a implemented by smtp.Auth of LOGIN mechanism
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", nil, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch {
		case bytes.Equal(fromServer, []byte("Username:")):
			return []byte(a.username), nil
		case bytes.Equal(fromServer, []byte("Password:")):
			return []byte(a.password), nil
		default:
			return nil, fmt.Errorf("unexpected server challenge: %s", fromServer)
		}
	}

	return nil, nil
}
