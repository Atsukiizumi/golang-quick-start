package define

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var (
	JwtKey = []byte("cloud-disk-key")
)
