package internal

import (
	"github.com/gin-gonic/gin"

	"CipherX/constant"
	"CipherX/internal/middleware"
	"CipherX/internal/router"
	res "CipherX/pkg/response"
)

func Routers() *gin.Engine {
	// Initialize the router
	Router := gin.New()

	// Middleware
	Router.Use(middleware.Logger())
	Router.Use(gin.Recovery())
	Router.Use(middleware.Install())

	PingGroup := Router.Group("")
	{
		// Ping
		PingGroup.GET("/ping", func(c *gin.Context) {
			res.ResSuccess(c, gin.H{
				"status":     "pong",
				"version":    constant.Version,
				"commit":     constant.Commit,
				"build_time": constant.BuildTime,
				"go_version": constant.GoVersion,
			})
		})
	}

	// Install
	InstallGroup := Router.Group("")
	router.InitRouterInstall(InstallGroup)

	return Router
}
