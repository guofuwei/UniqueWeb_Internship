package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"task4_2/middleware"

	"github.com/gin-gonic/gin"
)

type AuthMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.LoadHTMLGlob("templates/html/*")

	r.GET("thirdweb", thirdWeb)
	r.GET("", thirdWeb)

	r.GET("authstate", authStateHandle)
	r.GET("getlog2", getLog2Handle)

	r.Run("127.0.0.1:8001")

}

func thirdWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "thirdweb.html", gin.H{})
}

func authStateHandle(c *gin.Context) {
	code := c.Query("code")
	expiretime := c.Query("expiretime")
	request_url := "http://127.0.0.1:8000/oauth/token?client_id=TEST_ID&grant_type=authorization_code&code=" + code + "&redirect_uri=http:127.0.0.1:8001/thirdcatlog&expiretime=" + expiretime
	rsps, err := http.Post(request_url, "plain/text", nil)
	if err != nil {
		log.Fatal(err)
	}
	authMsg := &AuthMsg{}
	b, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, authMsg)
	if err != nil {
		log.Fatal(err)
	}
	time, _ := strconv.Atoi(expiretime)
	// fmt.Println("--------------------------------")
	// fmt.Println(time)
	c.SetCookie("authToken", authMsg.Data, 3600*time, "/", "", false, false)
	// fmt.Println(string(b))
	c.Redirect(http.StatusMovedPermanently, "getlog2")
}

func getLog2Handle(c *gin.Context) {
	c.HTML(http.StatusOK, "getlog2.html", gin.H{})
}
