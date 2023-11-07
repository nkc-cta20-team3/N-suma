package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// type UpdateAuthRequest struct {
// 	DocumentID     int    `json:"document_id"` //ドキュメントID
// 	UserID         int    `json:"user_id"`
// 	TeacherComment string `json:"teacher_comment"` //教員コメント
// }

// type UpdateDocument struct {
// 	Status         int    `json:"status"`          // ステータス
// 	TeacherComment string `json:"teacher_comment"` // 教員コメント
// }

func UpdateAuth(c *gin.Context) {

	request := model.UpdateAuthRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//必要な変数を定義
	var documentStatus int
	var post int
	message := ""

	//認可ステータスの取得
	db.Table("oa").Select("status").Where("document_id = ?", request.DocumentID).Scan(&documentStatus)

	//役職IDの取得
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	log.Print("リクエスト")
	log.Println(request)
	log.Print("認可する書類の状態")
	log.Println(documentStatus)
	log.Print("認可者の役職")
	log.Println(post)

	if post == 2 && documentStatus == 1 {
		//教員かつ認可できる書類の場合→認可処理
		db.Table("oa").
			Where("document_id = ?", request.DocumentID).
			Updates(model.UpdateDocument{Status: post})
		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}
	} else if post == 2 {
		//ステータスが異なるとき
		message = "DOCUMENT STATUS ERROR"
	} else if documentStatus == 1 {
		//認可者が不適切な時
		message = "POST ERROR"
	} else {
		//ステータスも認可者も違うとき
		message = "DOCUMENT STATUS & POST ERROR"
	}

	if message == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "SUCCESS",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
	}
}
