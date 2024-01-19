package infra

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

func init() {

	var err error

	// path := os.Getenv("DB_DSN")	// 接続文字列の取得
	// 接続文字列を作成する
	path := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=true"
	
	// 接続文字列を確認する
	// log.Println(os.Getenv("DB_DSN"))
	log.Println(path)
	
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
	DB.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
		SetConnMaxIdleTime(time.Minute * 30).
		SetConnMaxLifetime(time.Hour * 12).
		SetMaxIdleConns(40).
		SetMaxOpenConns(80),
	)
	// データベース接続成功
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