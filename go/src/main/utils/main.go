package utils

import (
	"log"
	"time"

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

// 文字列を時刻型に変換する関数
func StringToTime(stringTime string) time.Time {
	// 文字列を時刻型に変換
	time, err := time.Parse("2006-01-02 15:04", stringTime)
	// TODO: エラーハンドリング
	if err != nil {
		log.Println(err)
	}
	return time
}

// 時刻型を文字列に変換する関数
func TimeToString(time *time.Time) string {
	// 時刻型を文字列に変換
	stringTime := time.Format("2006-01-02 15:04")
	return stringTime
}

// 「2023-12-11T15:39:00.000Z」を「2023-12-11 15:39」に変換
func StringToTime2(stringTime string) time.Time {
	// 文字列を時刻型に変換
	time, err := time.Parse("2006-01-02T15:04:00.000Z", stringTime)
	if err != nil {
		log.Println(err)
	}
	return time
}

// 「2023-12-11T15:39:00.000Z」を「2023-12-11 15:39」に変換
func TimeToString2(time *time.Time) string {
	// 時刻型を文字列に変換
	stringTime := time.Format("2006-01-02 15:04")
	return stringTime
}

// 「2023-12-11T00:00:00Z」を「2023-12-11」に変換
func StringToTime3(stringTime string) time.Time {
	// 文字列を時刻型に変換
	time, err := time.Parse("2006-01-02T00:00:00Z", stringTime)
	if err != nil {
		log.Println(err)
	}
	return time
}

// 「2023-12-11T15:54:00Z」を「12-11 15:54」に変換
func StringToTime4(stringTime string) time.Time {
	// 文字列を時刻型に変換
	time, err := time.Parse("2006-01-02T15:04:00Z", stringTime)
	if err != nil {
		log.Println(err)
	}
	return time
}
