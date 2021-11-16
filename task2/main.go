package main

import (
	"fmt"
	"net/http"
	"time"

	"strconv"

	"task2/LRU"
	"task2/db_operate"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Binding from JSON
type data struct {
	Filename   string
	Expiretime string
	Code       string
}

func getpage(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", gin.H{
		"message": "",
	})
}

func postcode(c *gin.Context) {
	var json data
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    10001,
			"message": "The msg is no enough",
		})
		return
	}
	//开始检查filename是否重复
	ret := db_operate.Select(db_operate.Db)
	if _, ok := ret[json.Filename]; ok {
		c.JSON(http.StatusOK, gin.H{
			"code":    10002,
			"message": "该filename已存在",
		})
	}
	//对过期时间进行处理
	nowtime := time.Now().Unix()
	moretime, _ := strconv.Atoi(json.Expiretime)
	expiretime := nowtime + int64(moretime*60)
	//开始执行缓存操作
	LRU.LruCache.Set(json.Filename, json.Filename, expiretime, json.Code)
	//开始执行数据库操作
	db_operate.Insert(db_operate.Db, json.Filename, json.Code, expiretime)
	//返回消息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "",
	})
}

func getallcode(c *gin.Context) {
	ret := db_operate.Select(db_operate.Db)
	fmt.Println(ret)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": ret,
	})
}

func getcode(c *gin.Context) {
	filename := c.Param("filename")
	//先尝试从LRU缓存中拿取数据
	ret := LRU.LruCache.Get(filename)
	if ret != "" {
		fmt.Println("----cache in-----")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": ret,
		})
	} else {
		fmt.Println("----database in------")
		//当LRU缓存拿取失败时从数据库中查找
		ret := db_operate.SelectOne(db_operate.Db, filename)
		//将数据库查找到的内容缓存到LRU
		LRU.LruCache.Set(filename, filename, ret[filename].Expiretime, ret[filename].Code)
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": ret[filename].Code,
		})
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("post", getpage)
	r.POST("api/code", postcode)
	r.GET("getallcode", getallcode)
	r.GET("getcode/:filename", getcode)
	r.Run(":8000")
}

// fmt.Println(json)
// fmt.Printf("filename:%s\nexpiretime:%s\ncode:%s\n", json.Filename, json.Expiretime, json.Code)
