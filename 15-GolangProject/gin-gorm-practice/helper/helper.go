package helper

import (
	"GolangProject/gin-gorm-practice/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var myKey = []byte("gin-gorm-key")

func GetMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GetSaltMD5(s, salt string) string {
	p := "$" + salt + "$" + s
	return GetMD5(p)
}

// GenerateToken
// 生成 Token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:         identity,
		Name:             name,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	signedString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// AnalyseToken
// 解析 Token
func AnalyseToken(token string) (*UserClaims, error) {
	claims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	} else if claim, ok := claims.Claims.(*UserClaims); ok {
		fmt.Println(claim.Identity, claim.Name)
		return claim, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

// SendCode
// 发送验证码
func SendCode(toUserMail, code string) error {
	from := define.ServerUser
	pwd := define.ServerPwd
	smtpServer := define.ServerHost
	serverName := strings.Split(smtpServer, ":")[0]

	e := email.NewEmail()
	e.From = "Noreply <" + from + ">"
	e.To = []string{toUserMail}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码为： <b>" + code + "</b>")

	err := e.SendWithTLS(smtpServer,
		smtp.PlainAuth("", from, pwd, serverName),
		&tls.Config{InsecureSkipVerify: true, ServerName: serverName})

	return err
}

// SendCodeOMail
// 发送验证码，使用原生库
func SendCodeOMail(toUserMail, code string) error {
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
	defer dial.Close()

	conn, err := smtp.NewClient(dial, serverName)
	if err != nil {
		return fmt.Errorf("创建客户端失败：%v", err.Error())
	}
	defer conn.Quit()

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
	defer w.Close()

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

// GetUUID
// 生成 UUID
func GetUUID() string {
	return uuid.New().String()
}

// GetRandCode
// 生成验证码
func GetRandCode() string {
	dict := []byte("0123456789ABCDEFGHJKLMNPRTUVWXYZ")
	codeLen, _ := strconv.Atoi(define.DefaultCodeLen)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	s := ""
	rpt := make(map[string]int)
	for i := 0; i < codeLen; i++ {
		//设想如何不重复（流水码）
		char := string(dict[r.Intn(len(dict))])
		if rpt[char]++; rpt[char] > 1 {
			continue
		}
		s += char
	}
	return s
}
