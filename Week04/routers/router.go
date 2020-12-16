package routers

import (
	"github.com/gin-gonic/gin"
	"service-notification/pkg/middleware"
	"service-notification/pkg/setting"
	v1 "service-notification/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message" : "test",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.MiddlewareAuth())
	{
		//获取邮箱配置
		apiV1.POST("/mailbox/detail", v1.GetMailBox)
		//保存邮件配置
		apiV1.POST("/mailbox/save", v1.SaveMailbox)
		apiV1.POST("mailbox/test", v1.SendTestMail)
		//阿里云app推送API
		apiV1.POST("/alipush/accounts", v1.PushByAccounts)
		//短信api
		apiV1.POST("/sms/single_send", v1.SingleSend)
		apiV1.POST("/sms/callback", v1.XwCallback)
	}
	//ginpprof.Wrap(r)

	return r
}
