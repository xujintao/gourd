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
		log.Fatal("appid not exist")
	}
	appID, ok := v1.(string)
	if !ok {
		log.Fatal("appid is not string", appID)
	}

	// 获取usercode
	v2, ok := ctx.Get("usercode")
	if !ok {
		log.Fatal("usercode not exist")
	}
	userCode, ok := v2.(string)
	if !ok {
		log.Fatal("usercode is not string", userCode)
	}

	ctx.JSON(200, gin.H{
		"message":  "pong",
		"appid":    appID,
		"usercode": userCode,
	})
}
