package api

import (
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NextDocument(c *gin.Context) {
	//必要な変数(引数・戻り値)の定義
	request := []model.NextDocumentRequest{}
	response := []model.NextDocumentResponse{}

	var testUserID int
	var PostID int

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DBに接続
	db := infra.DBInitGorm()

	//ここでユーザIDと書類の整合性を確認(教員かどうか確認)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&PostID)

	if PostID == 1 {
		//担任用の処理
		//副問合せ用クエリ
		subquery := db.Table("user").
			Select("user.class_id").
			Joins("JOIN oa ON oa.user_id = user.user_id").
			Where("document_id = ?", request.DocumentID)

		//ドキュメントIDの学生を担当する教員のユーザIDを取得
		db.Table("user").Select("user_id,post_id").Where("class_id = ?", subquery).First(&testUserID)

		//担任かどうかをチェック
		if testUserID == request.UserID {

			//認可すべき次の書類を取得・格納する
			db.Table("oa").
				Select(
					"oa.document_id",
					"oa.request_at",
					"oa.start_time",
					"oa.start_flame",
					"oa.end_time",
					"oa.end_flame",
					"oa.location",
					"oa.student_comment",
					"oa.teacher_comment",
					"user.user_number",
					"cs.class_addr",
					"dv.division_name").
				Joins("JOIN user on oa.user_id = user.user_id").
				Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
				Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
				WHERE("user.user_id = ?", request.UserID).
				Where("oa.status =", PostID-1)
			Where()
			Order("oa.request_at ASC").
				First(&response)
			if db.Error != nil {
				fmt.Print("SQL ERROR!")
			}

			//戻り値
			c.JSON(http.StatusOK, gin.H{
				"message": response,
			})

		} else {
			//リクエストが不正及び誤りがあるとき
			c.JSON(http.StatusBadRequest, gin.H{"document": response})
		}
	}
}
