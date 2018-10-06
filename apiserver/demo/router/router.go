package router

import (
	"apiserver/demo/handler/sd"
	"apiserver/demo/handler/user"
	"apiserver/demo/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	/*
		gin.Recovery()：在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，
		这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	*/
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	//	404
	g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route.")
	})

	// api for authentication functionalities
	g.POST("/login", user.Login)

	/*
		服务需要进行分组
	*/
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// 健康检查
	svcd := g.Group("/sd")

	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
