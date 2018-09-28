package auth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// MyClaims 自定义Claims
type MyClaims struct {
	*jwt.StandardClaims
	AppID    int    `json:"appid,omitempty"`
	UserCode string `json:"usercode,omitempty"`
}

// JwtAuthMiddleware JWT认证中间件
func JwtAuthMiddleware(key string) gin.HandlerFunc {
	// claims := &MyClaims{} 	// claims放这里可能会有协程安全问题，先挪到下面

	return func(ctx *gin.Context) {
		claims := &MyClaims{}
		token, err := request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			// 先验证签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(key), nil
		}, request.WithClaims(claims))
		if err != nil {
			ctx.AbortWithError(401, err)
		}

		// 验证claims
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid && claims.Valid() == nil {
			// if claims.Valid() != nil {
			ctx.Set("appid", claims.AppID)
			ctx.Set("usercode", claims.UserCode)
			ctx.Next()
			return
		}
		ctx.AbortWithError(401, err)
	}
}
