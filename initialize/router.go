package initialize

import (
	"backend/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// Router.Use(cors.New(cors.Config{
	// 	AllowAllOrigins: true, // 允许所有源
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}, // 允许的HTTP方法
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},      // 允许的头部
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		PublicGroup.POST("command", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"status":  "Success",
				"message": "Command executed successfully",
			})
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	return Router
}
