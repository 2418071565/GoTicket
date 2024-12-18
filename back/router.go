package main

import (
	"github.com/2418071565/GoTicket/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(ser *gin.Engine) *gin.Engine {
	// 需要 token 鉴权的 API
	auth_group := ser.Group("/", controllers.JWTAuthMiddleware())
	{
		auth_user_group := auth_group.Group("")
		{
			auth_user_group.GET("/userInfo", controllers.GetUserInfo)
			auth_user_group.POST("/oneWayTickets", controllers.OneWayTicketsQuery)
			auth_user_group.POST("/route", controllers.AddRoute)
			auth_user_group.POST("/train", controllers.AddTrain)
			auth_user_group.POST("/station", controllers.AddStation)
			auth_user_group.POST("/trainNumber", controllers.AddTrainNumber)
		}
		auth_admin_group := auth_group.Group("/admin")
		{
			auth_admin_group.POST("", controllers.AdminRegister)
		}
	}

	// 普通 API
	ser.GET("/token", controllers.Login)
	ser.POST("/user", controllers.Register)
	return ser
}
