package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Pong for test
func Pong(ctx *gin.Context) {
	// 获取appid
	v1, ok := ctx.Get("appid")
	if !ok {
		log.Panic("appid not exist")
	}
	appID, ok := v1.(int)
	if !ok {
		log.Panic("appid is not int", appID)
	}

	// 获取usercode
	v2, ok := ctx.Get("usercode")
	if !ok {
		log.Panic("usercode not exist")
	}
	userCode, ok := v2.(string)
	if !ok {
		log.Panic("usercode is not string", userCode)
	}

	ctx.JSON(200, gin.H{
		"message":  "pong",
		"appid":    appID,
		"usercode": userCode,
	})
}
