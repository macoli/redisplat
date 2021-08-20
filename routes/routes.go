package routes

import (
	"net/http"

	"github.com/macoli/redisplat/controllers"

	"github.com/macoli/redisplat/logger"
	"github.com/macoli/redisplat/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//定义一个api 的路由组路由组
	v1 := r.Group("/api/v1")

	//v1.Use(middlewares.JWTAuthMiddleware()) //应用 JWT 认证中间件

	{
		// redis 监控相关路由
		v1.GET("/monitor/standalone", controllers.StandAlone)
		v1.GET("/monitor/sentinel", controllers.Sentinel)
		v1.GET("/monitor/cluster", controllers.Cluster)

	}

	//注册路由信息
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.AppConfig.Name)
	})

	return r
}
