package infra

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func DBInitGorm() *gorm.DB {
	// データベース接続
	db, err := gorm.Open(mysql.Open("root:password@tcp(db:3306)/cta20gr3?charset=utf8mb4&parseTime=true"))
	// データベース接続エラー
	if err != nil {
		log.Println(err)
	}

	return db
}