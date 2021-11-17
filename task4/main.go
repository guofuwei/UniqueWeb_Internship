package main

import (
	"task4/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*")
	routers.LoadAllRouters(r)
	r.Run(":8000")
}
