package routers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthManage struct {
	Name         string `json:"name"`
	WriteLog     string `json:"write_log"`
	CatOtherLog  string `json:"cat_other_log"`
	EditOtherLog string `json:"edit_other_log"`
}

func LoadAuthManage(e *gin.Engine) {
	e.GET("aumanage", auManageHandle)
	e.POST("api/aumanage", apiAuManageHandle)
}

func apiAuManageHandle(c *gin.Context) {
	var json AuthManage
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10001,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	sql := "update userinfo set write_log=?,cat_other_log=?,edit_other_log=? where username=?;"
	_, err = Db.Exec(sql, 2, "", "", json.WriteLog, json.CatOtherLog, json.EditOtherLog, json.Name)
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

func auManageHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "aumanage.html", gin.H{})
}
