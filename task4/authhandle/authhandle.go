package authhandle

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("This is a secret")

type LoginToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenLoginToken 生成登录Token
func GenLoginToken(username string) (string, error) {
	c := LoginToken{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "guofuwei",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseLoginToken 解析登录JWT
func ParseLoginToken(tokenString string) (*LoginToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &LoginToken{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*LoginToken); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
