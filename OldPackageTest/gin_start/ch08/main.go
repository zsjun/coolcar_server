package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	// 返回一个函数就可以了
	return func(c *gin.Context) {
		t := time.Now()
		// 添加内容到context中
		c.Set("example", "123456")
		// 让原本应该执行的逻辑继续执行
		c.Next()
		// 执行完以后，然后再执行之类的逻辑
		end := time.Since(t)
		fmt.Printf("耗时:%v\n", end)
		status := c.Writer.Status()
		fmt.Println("状态", status)
	}
}

func Hook404() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		status := c.Writer.Status()
		if status == 404 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "页面找不到",
			})
		}
	}
}

func main() {
	router := gin.Default()
	//使用logger和recovery中间件 全局所有
	router.Use(Hook404())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8083")
}
