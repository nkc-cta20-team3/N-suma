package infra

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBInitGorm() *gorm.DB {
	// データベース接続
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_HOST")), &gorm.Config{
		// ログ出力設定
		Logger: logger.Default.LogMode(logger.Info),
	})
	// データベース接続エラー
	if err != nil {
		log.Println(err)
	}
	return db
}
