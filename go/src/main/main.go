package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// ローカルモジュールのインポート
	"main/api"
	"main/apiHello"
)

func main() {

	//
	g := gin.Default()
	g.Use(cors.New(cors.Config{

		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
		},

		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"x-requested-with",
		},

		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,

		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// ルーティング
	routes := g.Group("/api")
	{
		// apiフォルダ内のapiをルーティング

		// curl -X POST http://localhost:8080/api/ral
		// curl -X POST -H "Content-Type: application/json" -d "{"user_id" : "1"}" http://localhost:8080/api/ral

		routes.POST("/cd", api.CreateDocument)    //公欠届作成
		routes.POST("/cu", api.CreateUser)        //ユーザ作成
		routes.POST("/nd", api.NextDocument)      //公欠届切り替え
		routes.POST("/al", api.ReadAlarm)         //通知取得
		routes.POST("/ral", api.ReadAuthList)     //未認可リスト取得
		routes.POST("/rd", api.ReadDocument)      //公欠届詳細取得
		routes.POST("/ra", api.RejectAuth)        //公欠届却下
		routes.POST("/rsd", api.ResubmitDocument) //公欠届再提出
		routes.POST("/ua", api.UpdateAuth)        //公欠届認可
		routes.POST("/uu", api.UpdateUser)        //ユーザ編集

		// 実装予定の管理者向けのAPI
		// routes.POST("/cd", api.CreateDocument)
		// routes.POST("/dd", api.DeleteDocument)
	}

	// ルーティング
	routes2 := g.Group("/hello")
	{	
		// 動作確認用
		// curl -X GET http://localhost:8080/hello
		routes2.GET("", apiHello.Hello)
		routes2.GET("/name", apiHello.Name)
		routes2.GET("/time", apiHello.Time)
		routes2.POST("/sum", apiHello.Sum)
	}

	g.Run(":8080")
}
