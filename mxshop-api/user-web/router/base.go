package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

// 注册用户相关的路由
func InitBaseRouter(Router *gin.RouterGroup){
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}

}
