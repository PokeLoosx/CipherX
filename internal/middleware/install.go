package middleware

import (
	"CipherX/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallRouter struct {
	url    string
	method string
}

const configFile = "config.yaml"

// WhiteList route list
var WhiteList = []InstallRouter{
	{
		url:    "/install",
		method: http.MethodGet,
	},
	{
		url:    "/install",
		method: http.MethodPost,
	},
	{
		url:    "/install/db/test",
		method: http.MethodPost,
	},
	{
		url:    "/install/redis/test",
		method: http.MethodPost,
	},
}

// Install checks if the configuration file exists
func Install() gin.HandlerFunc {
	return func(c *gin.Context) {
		exists := utils.PathFileExists(configFile)
		// Check if the current request is in the whitelist
		isWhite := false
		for _, route := range WhiteList {
			if c.Request.URL.Path == route.url && c.Request.Method == route.method {
				isWhite = true
				break
			}
		}

		// If the profile exists and the request is in the whitelist, access is forbidden
		if exists && isWhite {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "Forbidden access",
			})
			c.Abort()
			return
		}

		// If the configuration file does not exist and the request is not in the whitelist, it redirects to the installation page
		if !exists && !isWhite {
			c.Redirect(http.StatusFound, "/install")
			c.Abort()
			return
		}

		c.Next()
	}
}
