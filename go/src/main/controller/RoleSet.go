package controller

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type GetUser struct {
	UserID 		int 	`json:"user_id"` //ユーザID
	PostID 		int 	`json:"post_id"` //役職ID
	UserFlag  	bool  	`json:"user_flag"`
}

type CreateUser struct {
	UserUUID  	string 	`json:"user_uuid"`
	MailAddress string 	`json:"mail_address"`
	UserFlag  	bool   `json:"user_flag"`
}

func RoleSetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// コンテキストにセットされている、UUIDとEmailを取得する
		uuid := c.MustGet("UUID").(string)
		email := c.MustGet("Email").(string)

		// あらかじめコンテキストに-1をセットしておく
		c.Set("UserID", -1)
		c.Set("PostID", -1)

		// log.Printf("UserID: %v\n", c.MustGet("UserID").(int))
		// log.Printf("PostID: %v\n", c.MustGet("PostID").(int))
		
		errResponse := model.MessageError{}
		user := GetUser{}

		//DB接続とエラーハンドリング
		db := infra.DBInitGorm()
		if db.Error != nil {
			errResponse.Message = "データベース接続エラー"
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		
		// UUIDをもとに、DBからユーザー情報を取得する
		err := db.Table("user").
			Select("user_id, post_id, user_flag").
			Where("user_uuid = ?", uuid).
			Limit(1).
			Find(&user)
		if err.Error != nil {
			// その他のエラーハンドリング
			errResponse.Message = "OTHER ERROR"
			log.Println(err.Error.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		// ユーザー情報が存在しない場合は、ユーザーを作成する
		if err.RowsAffected == 0 {
			// 値のセット
			createUser := CreateUser{
				UserUUID: uuid,
				MailAddress: email,
				UserFlag: false,
			}
			
			// クエリの発行
			err := db.Table("user").
				Create(&createUser).
				Error
			if err != nil {
				// その他のエラーハンドリング
				errResponse.Message = "OTHER ERROR"
				log.Println(err.Error())
				c.JSON(http.StatusInternalServerError, errResponse)
				return
			}
		} else if user.UserFlag == true {
			// ユーザーが存在し、かつ、ユーザーが有効な場合

			// ロギング
			log.Println(user)
			// コンテキストに、ユーザーのIDをセットする
			c.Set("UserID", user.UserID)
			// コンテキストに、ユーザーのロールをセットする
			c.Set("PostID", user.PostID)
		}

		// 次のミドルウェアへ
		c.Next()
	}
}
