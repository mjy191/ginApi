package main

import (
	"ginApi/Routers"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.DisableConsoleColor()
	//file,_ := os.Create("access.log")
	//gin.DefaultWriter = io.MultiWriter(file,os.Stdout)
	r := gin.Default()
	r.Static("/static", "./static")
	//二级目录写法
	r.LoadHTMLGlob("template/**/*")
	Routers.Router{}.Router(r)
	r.Run(":8000")
}
