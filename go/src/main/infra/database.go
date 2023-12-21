package infra

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {

	var err error

	path := os.Getenv("DB_HOST")	// 接続文字列の取得
	dialector := mysql.Open(path)	// 接続文字列をdialectorに渡す
	// データベース接続
	DB, err = gorm.Open(dialector, &gorm.Config{
		// ログ出力設定
		Logger: logger.Default.LogMode(logger.Info),
	})
	// データベース接続リトライ
	if err != nil {
		log.Println(err)
		connect(dialector, 100)
	}
	log.Println("データベース接続成功")
}

// 初回のDB接続に失敗した時にリトライする
func connect(dialector gorm.Dialector, count uint) {
	var err error
	if DB, err = gorm.Open(dialector); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			log.Printf("retry... count:%v\n", count)
			connect(dialector, count)
			return
		}
		// データベース接続エラー
		log.Println("データベース接続エラー")
		panic(err.Error())
	}
}