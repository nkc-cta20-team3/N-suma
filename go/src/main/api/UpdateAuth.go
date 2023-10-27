package api

import (
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// type UpdateAuthRequest struct {
// 	DocumentID     int    `json:"document_id"`     //ドキュメントID
// 	UserNumber     int    `json:"user_number"`     //学内識別番号
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
	response := http.StatusBadRequest

	//認可ステータスの取得
	db.Table("oa").Select("status").Where("document_id = ?", request.DocumentID).Scan(&documentStatus)

	//役職IDの取得
	db.Table("user").Select("post_id").Where("user_number = ?", request.UserNumber).Scan(&post)

	log.Print("リクエスト")
	log.Println(request)
	log.Print("認可する書類の状態")
	log.Println(documentStatus)
	log.Print("認可者の役職")
	log.Println(post)

	//教員かどうかの判定
	if post == documentStatus+1 && documentStatus >= 0 {

		if post == 1 {

			//担任が認可するとき

			//認可書類の提出者と認可者が学生・担任関係にあるかの確認

			//更新処理
			db.Table("oa").
				Where("document_id = ?", request.DocumentID).
				Updates(model.UpdateDocument{Status: post, TeacherComment: request.TeacherComment})

		} else if post >= 2 {

			//教員(主任以上)が認可する

			//更新処理
			db.Table("oa").
				Where("document_id = ?", request.DocumentID).
				Updates(model.UpdateDocument{Status: post})

		}

		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		response = http.StatusOK

	} else {
		log.Print("認可書類が適切でない、または認可者が適切ではありません")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
