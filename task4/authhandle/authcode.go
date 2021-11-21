package authhandle

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthCode struct {
	WriteLog     string `json:"write_log"`
	CatOtherLog  string `json:"cat_other_log"`
	EditOtherLog string `json:"edit_other_log"`
	RedirectUri1 string `json:"redirect_uri1"`
	jwt.StandardClaims
}

func AuthCodeGen(arg map[string]string) (string, error) {
	writeLog := arg["writeLog"]
	catOtherLog := arg["catOtherLog"]
	editOtherLog := arg["editOtherLog"]
	redirectUri1 := arg["redirectUri1"]
	expiretime, _ := strconv.Atoi(arg["expiretime"])
	tokenExpireDuration := time.Hour * time.Duration(expiretime)
	c := AuthCode{
		writeLog,
		catOtherLog,
		editOtherLog,
		redirectUri1,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpireDuration).Unix(), // 过期时间
			Issuer:    "guofuwei",                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func ParseAuthCode(tokenString string) (*AuthCode, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &AuthCode{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthCode); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
