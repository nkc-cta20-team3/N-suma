package teacher

import (
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type ReadAlarmRequest struct {
// 	UserID int `json:"user_id"` //ユーザID
// }
// type TakePostID struct {
// 	PostID int `json:"post_id"` //役職ID
// }

func CheckAlarm(c *gin.Context) {

	//必要な変数定義
	request := model.CheckAlarmRequest{}
	take_post_id := model.TakePostID{}
	var AlarmFlag = false
	var count int64

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//役職を識別
	// db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Order("post_id DESC").First(&take_post_id)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Order("post_id DESC").First(&take_post_id)

	//ロールごとの処理分け

	if take_post_id.PostID == 2 {
		//教員→未認可リストの存在有無
		db.Table("oa").
			Select("document_id").
			Where("oa.status = ?", take_post_id.PostID-1).
			Count(&count)

		//ここから通知があったときの戻り値フラグをONにする
		if count > 0 {
			AlarmFlag = true
			fmt.Println("UNAUTH DOCUMENT EXIST")
		}

	} else {
		//エラー
		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}

	//結果を返却
	c.JSON(http.StatusOK, gin.H{"document": AlarmFlag})
}
