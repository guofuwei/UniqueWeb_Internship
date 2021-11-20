package routers

import (
	"log"
	"net/http"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

type Template1 struct {
	Username     string `json:"login_username"`
	TodaySummary string `json:"todaySummary"`
	TomorrowPlan string `json:"tomorrowPlan"`
	StudyTime    string `json:"studyTime"`
}
type Template2 struct {
	Username  string `json:"login_username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StudyTime string `json:"studyTime"`
}
type Template3 struct {
	Username        string `db:"login_username"`
	LastweekSummary string `json:"lastweekSummary"`
	ThisweekPlan    string `json:"thisweekPlan"`
}

func LoadTemplate(e *gin.Engine) {
	e.GET("template/1", template1Handle)
	e.GET("template/2", template2Handle)
	e.GET("template/3", template3Handle)
	e.POST("api/template/1", middleware.LoginMiddleware(), apiTemplate1Handle)
	e.POST("api/template/2", middleware.LoginMiddleware(), apiTemplate2Handle)
	e.POST("api/template/3", middleware.LoginMiddleware(), apiTemplate3Handle)
}

func template1Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "template1.html", gin.H{})
}

func template2Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "template2.html", gin.H{})
}
func template3Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "template3.html", gin.H{})
}

func apiTemplate1Handle(c *gin.Context) {
	var json Template1
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	//插入数据
	sql := "insert into loginfo(MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 1, json.TodaySummary, json.TomorrowPlan, json.StudyTime, "", "", "", "", json.Username)
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

func apiTemplate2Handle(c *gin.Context) {
	var json Template2
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	sql := "insert into loginfo(MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 2, "", "", json.StudyTime, json.Title, json.Content, "", "", json.Username)
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

func apiTemplate3Handle(c *gin.Context) {
	var json Template3
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	sql := "insert into loginfo(MyType,TodaySummary,TomorrowPlan,StudyTime,Title,Content,LastweekSummary,ThisweekPlan,Username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 3, "", "", "", "", "", json.LastweekSummary, json.ThisweekPlan, json.Username)
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
