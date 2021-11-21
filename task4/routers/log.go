package routers

import (
	"log"
	"net/http"
	"task4/authhandle"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

type LogInfo struct {
	Id              int    `db:"ID"`
	MyType          int    `db:"MyType"`
	TodaySummary    string `db:"TodaySummary"`
	TomorrowPlan    string `db:"TomorrowPlan"`
	StudyTime       int    `db:"StudyTime"`
	Title           string `db:"Title"`
	Content         string `db:"Content"`
	LastweekSummary string `db:"LastweekSummary"`
	ThisweekPlan    string `db:"ThisweekPlan"`
	Username        string `db:"Username"`
}

type LogUpdate1 struct {
	LogId        string `json:"log_id"`
	TodaySummary string `json:"todaySummary"`
	TomorrowPlan string `json:"tomorrowPlan"`
	StudyTime    string `json:"studyTime"`
}
type LogUpdate2 struct {
	LogId     string `json:"log_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StudyTime string `json:"studyTime"`
}
type LogUpdate3 struct {
	LogId           string `json:"log_id"`
	LastweekSummary string `json:"lastweekSummary"`
	ThisweekPlan    string `json:"thisweekPlan"`
}

func LoadLog(e *gin.Engine) {
	e.GET("getlog", getLogHandle)
	e.GET("api/getlog", middleware.LoginMiddleware(), apiGetLogHandle)
	e.GET("getlog2", getLogHandle2)
	e.GET("api/getlog2", apiGetLogHandle2)

	e.GET("api/getlog/edit1", editLog1Handle)
	e.GET("api/getlog/edit2", editLog2Handle)
	e.GET("api/getlog/edit3", editLog3Handle)
	e.GET("api/getlog/edit/data", editLogDataHandle)
	e.POST("api/log/update1", logUpdate1Handle)
	e.POST("api/log/update2", logUpdate2Handle)
	e.POST("api/log/update3", logUpdate3Handle)
}

func getLogHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "getlog.html", gin.H{})
}

func apiGetLogHandle(c *gin.Context) {
	// 先验证该用户有没有权限
	username := c.GetString("username")
	logAuth := authhandle.GetLogAuth(username)
	if !logAuth["catOtherLog"] {
		c.JSON(http.StatusOK, gin.H{
			"code": 10703,
			"msg":  "该用户没有权限查看",
			"data": "",
		})
	}
	logs := []LogInfo{}
	sqlstr := "select ID,MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username from loginfo;"
	err := Db.Select(&logs, sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": logs,
		"auth": logAuth,
	})
}

func getLogHandle2(c *gin.Context) {
	c.HTML(http.StatusOK, "getlog2.html", gin.H{})
}

func apiGetLogHandle2(c *gin.Context) {
	authToken := c.Query("authToken")
	// fmt.Println(authToken)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10801,
			"msg":  "非法请求",
			"data": "",
		})
		return
	}
	ret1, err := authhandle.ParseAuthToken(authToken)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10801,
			"msg":  "非法请求",
			"data": "",
		})
		return
	}
	// fmt.Println(ret1)
	ret2, err := authhandle.ParseAuthCode(ret1.Code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10801,
			"msg":  "非法请求",
			"data": "",
		})
		return
	}
	auth := make(map[string]string, 3)
	auth["writeLog"] = ret2.WriteLog
	auth["catOtherLog"] = ret2.CatOtherLog
	auth["editOtherLog"] = ret2.EditOtherLog
	//获取数据
	logs := []LogInfo{}
	sqlstr := "select ID,MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username from loginfo;"
	err = Db.Select(&logs, sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	//返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": logs,
		"auth": auth,
	})
}

func editLog1Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "editlog1.html", gin.H{})
}

func editLog2Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "editlog2.html", gin.H{})
}

func editLog3Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "editlog3.html", gin.H{})
}

func editLogDataHandle(c *gin.Context) {
	logId := c.Query("log_id")
	myLog := LogInfo{}
	sqlstr := "select ID,MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username from loginfo where ID=?;"
	err = Db.Get(&myLog, sqlstr, logId)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": myLog,
	})
}

func logUpdate1Handle(c *gin.Context) {
	var json LogUpdate1
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	//插入数据
	sql := "update loginfo set TodaySummary=?,TomorrowPlan=?,StudyTime=? where ID=?;"
	_, err = Db.Exec(sql, json.TodaySummary, json.TomorrowPlan, json.StudyTime, json.LogId)
	if err != nil {
		log.Fatal(err)
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": "",
	})
}

func logUpdate2Handle(c *gin.Context) {
	var json LogUpdate2
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	//插入数据
	sql := "update loginfo set Title=?,Content=?,StudyTime=? where ID=?;"
	_, err = Db.Exec(sql, json.Title, json.Content, json.StudyTime, json.LogId)
	if err != nil {
		log.Fatal(err)
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": "",
	})
}

func logUpdate3Handle(c *gin.Context) {
	var json LogUpdate3
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	//插入数据
	sql := "update loginfo set LastweekSummary=?,ThisweekPlan=? where ID=?;"
	_, err = Db.Exec(sql, json.LastweekSummary, json.ThisweekPlan, json.LogId)
	if err != nil {
		log.Fatal(err)
	}
	//返回
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": "",
	})
}
