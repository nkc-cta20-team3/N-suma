package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// gin-swaggerをインポート
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// docsのディレクトリを指定
	_ "main/docs"

	// ローカルモジュールのインポート

	// "main/api"
	apiHello "main/api/hello"
	apiAdminAdd "main/api/admin/add"
	apiAdminEdit "main/api/admin/edit"
	apiAdminList "main/api/admin/list"
	
	//"main/api/student"
	//"main/api/teacher"
	"main/controller"
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

	
	g := gin.Default()
	g.Use(controller.LogMiddleware())
	g.Use(cors.New(cors.Config{

		// アクセスを許可したいアクセス元
		AllowOrigins: []string{

			// ローカルのデプロイ先
			"http://localhost:5173",

			// awsのデプロイ先
			"http://n-suma.com",
			"https://n-suma.com",

			// firebaseのデプロイ先
			"http://n-suma.firebaseapp.com",
			"https://n-suma.firebaseapp.com",
			"http://n-suma.web.app",
			"https://n-suma.web.app",
			
		},

		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
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

	// swagger uiにアクセスするためのルーティング
	// swagger/index.html
	swaggerui := g.Group("/swagger")
	{
		swaggerui.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 導通兼用SwaggerGo用のサンプルコード
	hello := g.Group("/hello")
	{
		hello.GET("/get", apiHello.GetHello)
		hello.POST("/post", apiHello.PostHello)	
	}

	// 認証が必要なAPIのルーティング
	authRouter := g.Group("/")
	authRouter.Use(controller.AuthMiddleware())

	// 管理者用APIのルーティング
	admin := authRouter.Group("/admin")
	// admin.Use(controller.RoleMiddleware("Admin"))
	{
		admin.POST("/post", apiHello.PostHello)				// 導通確認用
	}
	
	// ユーザー登録画面用APIのルーティング
	adminAdd := admin.Group("/add")
	{
		adminAdd.POST("/read", apiAdminAdd.ReadPrepareInformation)		// 役職が未登録ユーザーのid,uuid,emailをリストで一覧取得
		adminAdd.POST("/create", apiAdminAdd.CreateUser)				// ユーザー作成
	}

	// ユーザー一覧画面用APIのルーティング
	adminList := admin.Group("/list")
	{
		adminList.POST("/read", apiAdminList.ReadUserList)	// ユーザー一覧取得
		adminList.POST("/sort", apiAdminList.SortUser)		// ユーザー検索
	}

	// ユーザー編集画面用APIのルーティング
	adminEdit := admin.Group("/edit")
	{
		adminEdit.POST("/read", apiAdminEdit.ReadUser)		// ユーザー詳細取得
		adminEdit.POST("/update", apiAdminEdit.UpdateUser)	// ユーザー編集
		adminEdit.POST("/delete", apiAdminEdit.DeleteUser)	// ユーザー削除
	}
	
	
	/*
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
	*/

	/*
	routes.POST("/cl", api.CheckLogin) //ログイン確認
	routes.POST("/rc", api.ReadClass)
	routes.POST("/rp", api.ReadPost)
	routes.POST("/rup", api.ReadUserPost)

	//調査用
	routes.POST("/td", api.TestDate)
	*/

	g.Run(":8080")
}
