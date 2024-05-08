package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"` //用户名
	Role     int8   `json:"role"`     // 1 管理员 2 普通用户
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GenToken  生成jwt token
func GenToken(JwtPayload JwtPayLoad, accessSecret string, expires int) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: JwtPayload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(accessSecret))
}

// ParseToken 解析 token
func ParseToken(tokenStr string, accessSecret string) (*CustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
