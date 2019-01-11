/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2019/1/11 23:06
 */
package utils

import (
	"fmt"
	"net/smtp"
	"noxue/config"
	"strings"
)

type loginAuth struct {
	username, password string
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}
func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}

func SendEmail(email, subject, content, contentType string) (err error){
	s:=config.Config.Smtp
	auth :=  &loginAuth{s.User, s.Pass}
	to := []string{email}
	if contentType==""{
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + s.Name +
		"<" + s.User + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + content)
	err = smtp.SendMail(fmt.Sprintf("%s:%d",s.Host,s.Port), auth, s.User, to, msg)
	return
}