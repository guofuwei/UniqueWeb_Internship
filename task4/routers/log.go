package routers

import (
	"log"
	"net/http"
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

func LoadLog(e *gin.Engine) {
	e.GET("getlog", getLogHandle)
	e.GET("api/getlog", middleware.LoginMiddleware(), apiGetLogHandle)
}

func getLogHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "getlog.html", gin.H{})
}

func apiGetLogHandle(c *gin.Context) {
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
	})
}
