package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	// 返回json
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {
	// 实例化一个gin的server对象
	r := gin.Default()
	// 定义一个接口
	r.GET("/ping", pong)
	// 监听8083端口
	r.Run(":8083") 
}
