package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"net/smtp"
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
	e := email.NewEmail()
	e.From = "Noreply <noreply@example.com>"
	e.To = []string{toUserMail}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码为： <b>" + code + "</b>")

	err := e.SendWithTLS("smtp.example.com:587",
		smtp.PlainAuth("", "", "", ""),
		&tls.Config{InsecureSkipVerify: true, ServerName: ""})

	return err
}

// GetUUID
// 生成 UUID
func GetUUID() string {
	return uuid.New().String()
}
