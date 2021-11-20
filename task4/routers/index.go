package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Loadindex(e *gin.Engine) {
	e.GET("/index", indexHandle)
	e.GET("", indexHandle)
}

func indexHandle(c *gin.Context) {
	//获取首页的信息
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
