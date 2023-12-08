package controller

import (
	"main/utils"
	"github.com/gin-gonic/gin"	
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ログ出力
		utils.LogClientIP(c)
		c.Next()
	}
}
