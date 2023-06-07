package infra

import (
	"log"

	"main/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func DBInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(db:3306)/cta20gr3?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	engine.ShowSQL(true)

	// ユーザーテーブルが存在するか確認
	exist, err := engine.IsTableExist("users")
	if err != nil {
		log.Fatal(err)
	}

	// 存在しなければテーブルを作成
	if !exist {
		engine.CreateTables(&model.Users{})
	}

	return engine
}
