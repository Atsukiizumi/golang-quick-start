package helper

import (
	"GolangProject/cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/smtp"
	"strings"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GeneratorToken(id int, identity, name string) (string, error) {
	userClaim := &define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	signing, err := token.SignedString(define.JwtKey) // 签名
	if err != nil {
		return "", err
	}
	return signing, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	// 解析token
	claims, err := jwt.ParseWithClaims(token, &define.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return define.JwtKey, nil
	})
	if err != nil {
		return nil, err
	} else if claim, ok := claims.Claims.(*define.UserClaim); ok {
		fmt.Println(claim.Identity, claim.Name)
		return claim, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

// SendCodeMail
// 发送验证码，使用原生库
func SendCodeMail(toUserMail, code string) error {
	from := define.ServerUser
	pwd := define.ServerPwd
	smtpServer := define.ServerHost
	serverName := strings.Split(smtpServer, ":")[0]

	auth := smtp.PlainAuth("", from, pwd, serverName)
	tslConfig := &tls.Config{InsecureSkipVerify: true, ServerName: serverName}

	dial, err := tls.Dial("tcp", smtpServer, tslConfig)
	if err != nil {
		return fmt.Errorf("TLS连接失败：%v", err.Error())
	}
	defer func(dial *tls.Conn) {
		err := dial.Close()
		if err != nil {
			log.Println("关闭连接失败：", err.Error())
		}
	}(dial)

	conn, err := smtp.NewClient(dial, serverName)
	if err != nil {
		return fmt.Errorf("创建客户端失败：%v", err.Error())
	}
	defer func(conn *smtp.Client) {
		err := conn.Quit()
		if err != nil {
			log.Println("关闭smtp失败：", err.Error())
		}
	}(conn)

	if err = conn.Auth(auth); err != nil {
		return fmt.Errorf("认证失败：%v", err.Error())
	}

	if err := conn.Mail(from); err != nil {
		return fmt.Errorf("发送邮件失败：%v" + err.Error())
	}
	if err := conn.Rcpt(toUserMail); err != nil {
		return fmt.Errorf("接收邮件失败：%v", err.Error())
	}

	w, err := conn.Data()
	if err != nil {
		return fmt.Errorf("发送数据失败：%v", err.Error())
	}
	defer func(w io.WriteCloser) {
		err := w.Close()
		if err != nil {
			log.Println("写入失败：", err.Error())
		}
	}(w)

	msg := []byte("From: " + from + "\r\n" +
		"To: " + toUserMail + "\r\n" +
		"Subject: 验证码\r\n" +
		"\r\n" +
		"您的验证码为： <b>" + code + " </b>。验证码有效时间为5分钟。" +
		"\r\n" +
		"本邮件为系统自动生成，请勿回复。\r\n")
	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("发送消息失败：%v", err.Error())
	}
	return nil
}
