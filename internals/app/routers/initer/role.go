package initer

import (
	"github.com/gin-gonic/gin"
	"rmonitor/internals/app/middleware"
	v1 "rmonitor/internals/app/routers/v1"
)

func RoleRouterInit(roleRouterGroup *gin.RouterGroup) {
	roleRouterGroup.Use(middleware.JWTAuth())
	{
		roleRouterGroup.GET("/list", v1.RoleList)
		roleRouterGroup.POST("/add_one", v1.RoleAddOne)
		roleRouterGroup.POST("/del_one", v1.RoleDeleteFromName)
		roleRouterGroup.POST("/update_one", v1.RoleUpdateOneFromName)
	}
}
