package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Template1 struct {
	username     string
	todaySummary string
	tomorrowPlan string
	studyTime    string
}
type Template2 struct {
	username  string
	title     string
	content   string
	studyTime string
}
type Template3 struct {
	username        string
	lastweekSummary string
	thisweekPlan    string
}

func LoadTemplate(e *gin.Engine) {
	e.GET("/template/1", template1Handle)
	e.GET("/template/2", template2Handle)
	e.GET("/template/3", template3Handle)
}

func template1Handle(c *gin.Context) {
	var json Template1
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	//插入数据
	sql := "insert into loginfo(type,todaySummary,tomorrowPlan,studyTime,title,content,lastweekSummary,thisweekPlan,user_username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 1, json.todaySummary, json.tomorrowPlan, json.studyTime, "", "", "", "", json.username)
	if err != nil {
		log.Fatal(err)
	}
}

func template2Handle(c *gin.Context) {
	var json Template2
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	sql := "insert into loginfo(type,todaySummary,tomorrowPlan,studyTime,title,content,lastweekSummary,thisweekPlan,user_username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 1, "", "", json.studyTime, json.title, json.content, "", "", json.username)
	if err != nil {
		log.Fatal(err)
	}
}

func template3Handle(c *gin.Context) {
	var json Template3
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	sql := "insert into loginfo(type,todaySummary,tomorrowPlan,studyTime,title,content,lastweekSummary,thisweekPlan,user_username) values(?,?,?,?,?,?,?,?,?);"
	_, err = Db.Exec(sql, 1, "", "", "", "", "", json.lastweekSummary, json.thisweekPlan, json.username)
	if err != nil {
		log.Fatal(err)
	}
}
