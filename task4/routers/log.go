package routers

import "github.com/gin-gonic/gin"

func LoadLog(e *gin.Engine) {
	e.GET("/getlog", getlogHandle)
}

func getlogHandle(c *gin.Context) {

}
