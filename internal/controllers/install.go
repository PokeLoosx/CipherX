package controllers

import (
	"net/http"

	"CipherX/internal/fields"
	"CipherX/internal/service"
	res "CipherX/pkg/response"

	"github.com/gin-gonic/gin"
)

type InstallController struct{}

// Router Api
func (c *InstallController) Router(r *gin.RouterGroup) {
	r.POST("install", c.Install)
	r.POST("install/db/test", c.InstallDBTest)
	r.POST("install/redis/test", c.InstallRedisTest)
}

// Install Generating Configuration Files
func (c *InstallController) Install(ctx *gin.Context) {
	var field fields.InstallFields
	if err := ctx.ShouldBindJSON(&field); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resCode, msg := service.Install(field)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // Success
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // Failed
	}
}

// InstallDBTest Testing database connectivity
func (c *InstallController) InstallDBTest(ctx *gin.Context) {
	var field fields.DBTestFields
	if err := ctx.ShouldBindJSON(&field); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resCode, msg := service.InstallDBTest(field)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // Success
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // Failed
	}
}

// InstallRedisTest Testing redis connectivity
func (c *InstallController) InstallRedisTest(ctx *gin.Context) {
	var field fields.RedisTestFields
	if err := ctx.ShouldBindJSON(&field); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resCode, msg := service.InstallRedisTest(field)
	if resCode == res.CodeSuccess {
		res.ResSuccess(ctx, msg) // Success
	} else {
		res.ResErrorWithMsg(ctx, resCode, msg) // Failed
	}
}
