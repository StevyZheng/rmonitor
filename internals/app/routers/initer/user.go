package initer

import (
	"github.com/gin-gonic/gin"
	"rmonitor/internals/app/middleware"
	v1 "rmonitor/internals/app/routers/v1"
)

func UserRouterInit(userRouterGroup *gin.RouterGroup) {
	userRouterGroup.Use(middleware.JWTAuth())
	{
		userRouterGroup.GET("/list", v1.UserList)
		userRouterGroup.POST("/add_one", v1.UserAddOne)
		userRouterGroup.POST("/del_one", v1.UserDeleteFromName)
		userRouterGroup.POST("/update_one", v1.UserUpdateOneFromName)
	}
}
