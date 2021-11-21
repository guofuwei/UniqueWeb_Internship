package routers

import (
	"log"
	"net/http"
	"task4/authhandle"
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

type CodePostStruct struct {
	WriteLog     string `json:"write_log"`
	CatOtherLog  string `json:"cat_other_log"`
	EditOtherLog string `json:"edit_other_log"`
	RedirectUri1 string `json:"redirect_uri1"`
	Expiretime   string `json:"expiretime"`
}

func LoadOauth(e *gin.Engine) {
	e.GET("oauth/code", oauthCodeHandle)
	e.POST("oauth/token", oauthTokenHandle)
	e.GET("api/oauth/code", middleware.LoginMiddleware(), apiOauthCodeHandle)
	e.POST("api/oauth/code/post", apiOauthCodePostHandle)
}

func oauthCodeHandle(c *gin.Context) {
	// responseType := c.Query("response_type")
	// clientId := c.Query("client_id")
	// redirectUrl := c.Query("redirect_uri")
	// data := make(map[string]string, 1)
	// data["redirectUrl1"] = redirectUrl
	// if responseType != "code" || clientId != "TES_ID" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 10601,
	// 		"msg":  "非法请求",
	// 		"data": data,
	// 	})
	// }
	c.HTML(http.StatusOK, "oauth.html", gin.H{})
}

func oauthTokenHandle(c *gin.Context) {
	clientId := c.Query("client_id")
	grantType := c.Query("grant_type")
	code := c.Query("code")
	redirectUrl := c.Query("redirect_uri")
	expiretime := c.Query("expiretime")
	arg := make(map[string]string, 4)
	arg["clientId"] = clientId
	arg["grantType"] = grantType
	arg["code"] = code
	arg["redirectUrl"] = redirectUrl
	arg["expiretime"] = expiretime
	authToken, err := authhandle.GenAuthToken(arg)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": authToken,
	})
}

func apiOauthCodeHandle(c *gin.Context) {
	username := c.GetString("username")
	logAuth := authhandle.GetLogAuth(username)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": logAuth,
	})
}

func apiOauthCodePostHandle(c *gin.Context) {
	var json CodePostStruct
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10701,
			"msg":  "The msg is no enough",
			"data": "",
		})
		return
	}
	arg := make(map[string]string, 5)
	arg["writeLog"] = json.WriteLog
	arg["catOtherLog"] = json.CatOtherLog
	arg["editOtherLog"] = json.EditOtherLog
	arg["expiretime"] = json.Expiretime
	arg["redirectUri1"] = json.RedirectUri1
	authCode, err := authhandle.AuthCodeGen(arg)
	if err != nil {
		log.Fatal(err)
	}
	redictUrl := json.RedirectUri1 + "?code=" + authCode + "&expiretime=" + json.Expiretime
	// _, err = http.Post(redictUrl, "plain/text", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := make(map[string]string, 1)
	data["redictUrl"] = redictUrl
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
	// fmt.Println(redictUrl)
	// c.Redirect(302, redictUrl)
	// c.Abort()
}
