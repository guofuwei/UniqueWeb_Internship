package routers

import "github.com/gin-gonic/gin"

func LoadAllRouters(e *gin.Engine) {
	LoadUser(e)
	LoadLog(e)
	LoadTemplate(e)
	Loadindex(e)
	LoadOauth(e)
	LoadAuthManage(e)
}
