package api

import (
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type UpdateDocument struct {
	Status         int    `json:"status"`          // ステータス
	TeacherComment string `json:"teacher_comment"` // 教員コメント
}

// RejectAuthで使用する構造体
type RejectAuthRequest struct {
	DocumentID     int    `json:"document_id"`
	UserID         int    `json:"user_id"`
	TeacherComment string `json:"teacher_comment"` // 教員コメント
}

func RejectAuth(c *gin.Context) {

	request := model.RejectAuthRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.TeacherComment == "" {
		//値が消失していた時、または何も入っていない時
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "TEACHER COMMENT NOT EXISTS",
		})
	} else {
		//DB接続
		db := infra.DBInitGorm()
		//却下
		db.Table("oa").
			Where("document_id = ?", request.DocumentID).
			Updates(model.UpdateDocument{Status: -1, TeacherComment: request.TeacherComment})

		//エラーハンドリング
		if db.Error != nil {
			errMsg := "データベース接続エラー"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": http.StatusOK,
		})
	}

}
