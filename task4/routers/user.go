package routers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type RegisterData struct {
	username  string
	password1 string
	password2 string
}

type LoginData struct {
	username string
	password string
}

type UserInfo struct {
	username     string
	password     string
	identify     string
	writeLog     bool
	catOtherLog  bool
	editOtherLog bool
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
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandle)
}

func loginHandler(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	// 开始处理参数
	m5 := md5.New()
	m5.Write([]byte(json.password))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	user := UserInfo{}
	Db.Get(&user, "select * from userinfo where username=?;", json.username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10021,
			"msg":  "用户名或密码错误",
			"data": "",
		})
		return
	}
	if pwdHash != user.password {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10022,
			"msg":  "用户名或密码错误",
			"data": "",
		})
		return
	}
	//密码正确
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": "", // 注意这个地方应该返回token
	})
}

func registerHandle(c *gin.Context) {
	var json RegisterData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	if json.password1 != json.password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10011,
			"msg":  "两次密码输入不一致",
			"data": "",
		})
		return
	}
	// 开始注册账户
	user := UserInfo{}
	err = Db.Get(&user, "select * from userinfo where username=?", json.username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10012,
			"msg":  "该账户已注册",
			"data": "",
		})
		return
	}
	m5 := md5.New()
	m5.Write([]byte(json.password1))
	pwdHash := hex.EncodeToString(m5.Sum(nil))
	sql := "insert into userinfo values(?,?,?,?,?,?);"
	_, err = Db.Exec(sql, json.username, pwdHash, "group member", 1, 1, 0)
	if err != nil {
		log.Fatal(err)
	}
}
