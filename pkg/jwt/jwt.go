package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TokenExpireDuration token的过期时间
const TokenExpireDuration = time.Hour * 100

//Payload jwt的payload部分
type Payload struct {
	Username string `json:"username"`

	ExpiresAt int64 `json:"exp"`
}

//Valid 效验
func (p Payload) Valid() error {
	return nil
}

// GenToken 生成JWT
func GenToken(username string, screat string) (string, error) {
	payload := Payload{
		Username:  username,
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(screat))
}
