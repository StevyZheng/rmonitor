package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rmonitor/internals/app/routers/initer"
	v1 "rmonitor/internals/app/routers/v1"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"app":     "rmonitor",
			"version": "1.0",
		})
	})

	apiV1 := router.Group("/api/v1")
	apiV1.POST("/login", v1.Login)

	apiRole := apiV1.Group("/role")
	initer.RoleRouterInit(apiRole)
	apiUser := apiV1.Group("/user")
	initer.UserRouterInit(apiUser)

	return router
}
