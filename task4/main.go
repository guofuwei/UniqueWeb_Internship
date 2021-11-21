package main

import (
	"task4/middleware"
	"task4/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.LoadHTMLGlob("templates/html/*")
	routers.LoadAllRouters(r)
	r.Run("127.0.0.1:8000")
}
