package infra

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInitGorm() *gorm.DB {
	// データベース接続
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_HOST")))
	// データベース接続エラー
	if err != nil {
		log.Println(err)
	}

	return db
}
