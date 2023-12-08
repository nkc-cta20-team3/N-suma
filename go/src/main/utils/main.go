package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// ログ出力を行う
func LogOutput(message string) {
	log.Println(message)
}

// アクセス元のIPアドレスを出力する
func LogClientIP(c *gin.Context) {
	log.Printf("ClientIP: %s\n", c.ClientIP())
}