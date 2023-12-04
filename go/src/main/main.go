package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// gin-swaggerをインポート
	ginSwagger "github.com/swaggo/gin-swagger" 
	swaggerFiles "github.com/swaggo/files" 
	// docsのディレクトリを指定
	_ "main/docs"
	
	// ローカルモジュールのインポート

	"main/api"
	"main/api/admin"
	"main/api/student"
	"main/api/teacher"
	"main/apiHello"
)


// @title N-suma-API
// @version 1.0.0
// @description N-suma用にGo側で実装するAPI
// @contact name:4年A組3班 Nスマ開発チーム 
// @contact url: 'https://github.com/nkc-cta20-team3/N-suma' 
// @contact email:dev@example
// @license name:MIT
// @license url: 'https://opensource.org/licenses/MIT'
// @host localhost:8080
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

		/*
			routes.POST("/cd", api.CreateDocument)    //公欠届作成
			routes.POST("/cu", api.CreateUser)        //ユーザ作成
			routes.POST("/nd", api.NextDocument)      //公欠届切り替え
			routes.POST("/ca", api.CheckAlarm)        //通知取得
			routes.POST("/ral", api.ReadAuthList)     //未認可リスト取得
			routes.POST("/rd", api.ReadDocument)      //公欠届詳細取得
			routes.POST("/ra", api.RejectAuth)        //公欠届却下
			routes.POST("/rsd", api.ResubmitDocument) //公欠届再提出
			routes.POST("/ua", api.UpdateAuth)        //公欠届認可
			routes.POST("/uu", api.UpdateUser)        //ユーザ編集

			routes.POST("/cl", api.CheckLogin)              //ログイン確認
			routes.POST("/rdv", api.ReadDivision)           //区分取得
			routes.POST("/al", api.ReadAlarm)               //通知詳細取得
			routes.POST("/rul", api.ReadUserList)           //ユーザリスト取得
			routes.POST("/ru", api.ReadUser)                //ユーザ詳細取得
			routes.POST("/su", api.SortUser)                //ユーザソート
			routes.POST("/du", api.DeleteUser)              //ユーザ削除
			routes.POST("/rpi", api.ReadPrepareInformation) //登録用情報取得
		*/
		routes.POST("/cl", api.CheckLogin) //ログイン確認
	}

  // SwaggerGo用のサンプルコード
	hello := g.Group("/hello")
	{	
		// 動作確認用
		// curl -X GET http://localhost:8080/hello
		hello.GET("", apiHello.Hello)
		hello.GET("/name", apiHello.Name)
		hello.GET("/time", apiHello.Time)
		hello.POST("/sum", apiHello.Sum)
	}

	
	// swagger uiにアクセスするためのルーティング
	swaggerui := g.Group("/swagger")
	{	
		swaggerui.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//学生がアクセスできるAPI
	routes3 := g.Group("/api/student")
	{

		//学生用
		routes3.POST("/nd", student.NextDocument)      //公欠届切り替え
		routes3.POST("/ral", student.ReadAuthList)     //未認可リスト取得
		routes3.POST("/rd", student.ReadDocument)      //公欠届詳細取得
		routes3.POST("/cd", student.CreateDocument)    //公欠届作成
		routes3.POST("/rsd", student.ResubmitDocument) //公欠届再提出
		routes3.POST("/ca", student.CheckAlarm)        //通知取得
		routes3.POST("/rdv", student.ReadDivision)     //区分取得
		routes3.POST("/al", student.ReadAlarm)         //通知詳細取得
	}

	routes4 := g.Group("/api/teacher")
	{
		routes4.POST("/nd", teacher.NextDocument)  //公欠届切り替え
		routes4.POST("/ral", teacher.ReadAuthList) //未認可リスト取得
		routes4.POST("/rd", teacher.ReadDocument)  //公欠届詳細取得
		routes4.POST("/ua", teacher.UpdateAuth)    //公欠届認可
		routes4.POST("/ra", teacher.RejectAuth)    //公欠届却下
		routes4.POST("/ca", teacher.CheckAlarm)    //通知取得
		routes4.POST("/al", teacher.ReadAlarm)     //通知詳細取得
	}

	routes5 := g.Group("/api/admin")
	{
		routes5.POST("/cu", admin.CreateUser)              //ユーザ作成
		routes5.POST("/uu", admin.UpdateUser)              //ユーザ編集
		routes5.POST("/rul", admin.ReadUserList)           //ユーザリスト取得
		routes5.POST("/ru", admin.ReadUser)                //ユーザ詳細取得
		routes5.POST("/su", admin.SortUser)                //ユーザソート
		routes5.POST("/du", admin.DeleteUser)              //ユーザ削除
		routes5.POST("/rpi", admin.ReadPrepareInformation) //登録用情報取得

	}

	g.Run(":8080")
}
