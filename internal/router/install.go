package router

import (
	"CipherX/internal/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouterOpen Api
func InitRouterInstall(r *gin.RouterGroup) {
	install := controllers.InstallController{}
	install.Router(r)
}
