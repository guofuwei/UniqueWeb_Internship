package routers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"task4/authhandle"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type RegisterData struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Username     string `db:"username"`
	Password     string `db:"password"`
	Identify     string `db:"identify"`
	WriteLog     bool   `db:"write_log"`
	CatOtherLog  bool   `db:"cat_other_log"`
	EditOtherLog bool   `db:"edit_other_log"`
}

var Db *sqlx.DB
var err error

func init() {
	Db, err = sqlx.Open("mysql", "task:123456@tcp(127.0.0.1:3306)/task")
	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadUser(e *gin.Engine) {
	e.GET("login", loginHandler)
	e.GET("register", registerHandle)
	e.POST("api/login", apiLoginHandler)
	e.POST("api/register", apiRegisterHandle)
}

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func registerHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func apiLoginHandler(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	// 开始处理参数
	m5 := md5.New()
	m5.Write([]byte(json.Password))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	user := UserInfo{}
	Db.Get(&user, "select * from userinfo where username=?;", json.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10021,
			"msg":  "用户名或密码错误",
			"data": "",
		})
		return
	}
	if pwdHash != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"code": 10022,
			"msg":  "用户名或密码错误",
			"data": "",
		})
		return
	}
	// 密码正确
	// 开始生成Token
	login_token, err := authhandle.GenLoginToken(json.Username)
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[string]string, 2)
	data["username"] = json.Username
	data["token"] = login_token
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data, // 注意这个地方应该返回token
	})
}

func apiRegisterHandle(c *gin.Context) {
	var json RegisterData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	if json.Password1 != json.Password2 {
		c.JSON(http.StatusOK, gin.H{
			"code": 10011,
			"msg":  "两次密码输入不一致",
			"data": "",
		})
		return
	}
	// 开始注册账户
	// user := UserInfo{}
	// err = Db.Get(&user, "select * from userinfo where username=?;", json.Username)
	// if err == nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 10012,
	// 		"msg":  "该账户已注册",
	// 		"data": "",
	// 	})
	// 	return
	// }
	// if user.username != "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 10012,
	// 		"msg":  "该账户已注册",
	// 		"data": "",
	// 	})
	// 	return
	// }
	m5 := md5.New()
	m5.Write([]byte(json.Password1))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	sql := "insert into userinfo values(?,?,?,?,?,?);"
	_, err = Db.Exec(sql, json.Username, pwdHash, "group member", 1, 1, 0)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10012,
			"msg":  "该账户已注册",
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": "",
	})
}
