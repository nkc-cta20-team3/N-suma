package controller

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	UserID int `json:"user_id"` //ユーザID
	PostID int `json:"post_id"` //役職ID
}

type CreateUser struct {
	UserUUID  	string 	`json:"user_uuid"`
	MailAddress string 	`json:"mail_address"`
}

func RoleSetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// コンテキストにセットされている、UUIDとEmailを取得する
		uuid := c.MustGet("UUID").(string)
		email := c.MustGet("Email").(string)

		// あらかじめコンテキストにnilをセットしておく
		c.Set("UserID", nil)
		c.Set("PostID", nil)

		errResponse := model.MessageError{}
		response := GetUserResponse{}

		//DB接続とエラーハンドリング
		db := infra.DBInitGorm()
		if db.Error != nil {
			errResponse.Message = "データベース接続エラー"
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		
		// UUIDをもとに、DBからユーザー情報を取得する
		err := db.Table("user").
			Select("user_id, post_id").
			Where("user_uuid = ?", uuid).
			Limit(1).
			Find(&response)
		if err.Error != nil {
			// その他のエラーハンドリング
			errResponse.Message = "OTHER ERROR"
			fmt.Println(err.Error.Error())
			c.JSON(http.StatusInternalServerError, errResponse)
			return
		}
		// ユーザー情報が存在しない場合は、ユーザーを作成する
		if err.RowsAffected == 0 {
			// 値のセット
			user := CreateUser{
				UserUUID: uuid,
				MailAddress: email,
			}
			
			// クエリの発行
			err := db.Table("user").
				Create(&user).
				Error
			if err != nil {
				// その他のエラーハンドリング
				errResponse.Message = "OTHER ERROR"
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, errResponse)
				return
			}
		} else {
			// ロギング
			fmt.Println(response)
			// コンテキストに、ユーザーのIDをセットする
			c.Set("UserID", response.UserID)
			// コンテキストに、ユーザーのロールをセットする
			c.Set("PostID", response.PostID)
		}

		// 次のミドルウェアへ
		c.Next()
	}
}
