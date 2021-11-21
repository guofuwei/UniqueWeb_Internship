package authhandle

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	ClientId     string `json:"clientId"`
	GrantType    string `json:"grantType"`
	Code         string `json:"code"`
	RedirectUri2 string `json:"redirectUrl"`
	jwt.StandardClaims
}

func GenAuthToken(arg map[string]string) (string, error) {
	clientId := arg["clientId"]
	grantType := arg["grantType"]
	code := arg["code"]
	redirectUrl := arg["redirectUrl"]
	expiretime, _ := strconv.Atoi(arg["expiretime"])
	tokenExpireDuration := time.Hour * time.Duration(expiretime)
	c := AuthToken{
		clientId,
		grantType,
		code,
		redirectUrl,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(), // 过期时间
			Issuer:    "guofuwei",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseAuthToken 解析登录JWT
func ParseAuthToken(tokenString string) (*AuthToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &AuthToken{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthToken); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
