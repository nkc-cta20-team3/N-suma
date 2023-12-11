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
	apiUtils "main/api/utils"

	apiAdminAdd "main/api/admin/add"
	apiAdminEdit "main/api/admin/edit"
	apiAdminList "main/api/admin/list"
	
	apiStudentAlarm "main/api/student/alarm"
	apiStudentForm "main/api/student/form"
	apiStudentReForm "main/api/student/reform"
	apiStudentList "main/api/student/list"
	apiStudentView "main/api/student/view"

	apiTeacherAlarm "main/api/teacher/alarm"
	apiTeacherUnAuthList "main/api/teacher/unauthlist"
	apiTeacherAuth "main/api/teacher/auth"
	apiTeacherViewList "main/api/teacher/viewlist"
	apiTeacherView "main/api/teacher/view"
	
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

	// 認証不要なAPIのルーティング
	router := g.Group("/utils")
	{
		router.GET("/read/class", apiUtils.ReadClass)		// クラス一覧取得
		router.GET("/read/post", apiUtils.ReadPost)			// 役職一覧取得
		router.GET("/read/division", apiUtils.ReadDivision)	// ユーザー役職一覧取得
	}

	// 認証が必要なAPIのルーティング
	authRouter := g.Group("/")
	authRouter.Use(controller.AuthMiddleware())

	// 管理者用APIのルーティング
	admin := authRouter.Group("/admin")
	// admin.Use(controller.RoleMiddleware("Admin"))
	
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
		adminList.POST("/search", apiAdminList.SearchUserList)		// ユーザー検索
	}

	// ユーザー編集画面用APIのルーティング
	adminEdit := admin.Group("/edit")
	{
		adminEdit.POST("/read", apiAdminEdit.ReadUser)		// ユーザー詳細取得
		adminEdit.POST("/update", apiAdminEdit.UpdateUser)	// ユーザー編集
		adminEdit.POST("/delete", apiAdminEdit.DeleteUser)	// ユーザー削除
	}
	
	// 学生用のAPIのルーティング
	student := authRouter.Group("/student")
	// student.Use(controller.RoleMiddleware("student"))
	
	// 通知関連のAPIのルーティング
	studentAlarm := student.Group("/alarm")
	{
		student.POST("/check", apiStudentAlarm.CheckAlarm)		// 通知が存在するかの確認
		student.POST("/read", apiStudentAlarm.ReadAlarm)		// 通知の内容を取得
	}

	// 書類提出画面用APIのルーティング
	studentForm := student.Group("/form")
	{		
		studentForm.POST("/create", apiStudentForm.CreateDocument)	// 書類提出
	}

	// 書類再提出画面用APIのルーティング
	studentReForm := student.Group("/reform")
	{
		studentReForm.POST("/create", apiStudentReForm.ReSubmitDocument)	// 書類再提出
	}

	// 提出書類一覧画面用APIのルーティング
	studentList := student.Group("/list")
	{
		studentList.POST("/read", apiStudentList.ReadDocumentList)	// 書類提出履歴一覧取得
	}

	// 書類提出履歴画面用APIのルーティング
	studentView := student.Group("/view")
	{
		studentView.POST("/read", apiStudentView.ReadDocument)		// 書類詳細取得
		studentView.POST("/next", apiStudentView.NextDocument)		// 書類切り替え
	}


	// 教員用のAPIのルーティング
	teacher := authRouter.Group("/teacher")
	// teacher.Use(controller.RoleMiddleware("teacher"))
	
	// 通知関連のAPIのルーティング
	teacherAlarm := teacher.Group("/alarm")
	{
		teacher.POST("/check", apiTeacherAlarm.CheckAlarm)		// 通知が存在するかの確認
		teacher.POST("/read", apiTeacherAlarm.ReadAlarm)		// 通知の内容を取得
	}

	// 書類認可画面用APIのルーティング
	teacherAuth := teacher.Group("/auth")
	{
		teacherAuth.POST("/read", apiTeacherAuth.ReadDocument)	// 書類詳細取得
		teacherAuth.POST("/reject", apiTeacherAuth.RejectAuth)	// 書類却下
		teacherAuth.POST("/update", apiTeacherAuth.UpdateAuth)	// 書類認可
	}

	// 未認可書類一覧画面用APIのルーティング
	teacherUnAuthLlist := teacher.Group("/unauthllist")
	{
		teacherUnAuthLlist.POST("/read", apiTeacherUnAuthList.ReadUnAuthList)	// 未認可書類一覧取得
	}	

	// 書類詳細画面用APIのルーティング
	teacherView := teacher.Group("/view")
	{
		teacherView.POST("/read", apiTeacherView.ReadDocument)	// 書類詳細取得
		teacherView.POST("/next", apiTeacherView.NextDocument)	// 書類切り替え
	}
	
	// 認可済書類一覧画面用APIのルーティング
	teacherViewList := teacher.Group("/viewlist")
	{
		teacherViewList.POST("/read", apiTeacherViewList.ReadDocumentList)	// 認可済書類一覧取得
	}
	
	/*
	routes.POST("/cl", api.CheckLogin) //ログイン確認
	routes.POST("/rup", api.ReadUserPost)

	//調査用
	routes.POST("/td", api.TestDate)
	*/

	g.Run(":8080")
}
