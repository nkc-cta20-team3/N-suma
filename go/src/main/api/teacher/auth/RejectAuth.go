package teacher

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

// type UpdateDocument struct {
// 	Status         int    `json:"status"`          // ステータス
// 	TeacherComment string `json:"teacher_comment"` // 教員コメント
// }

// // RejectAuthで使用する構造体
//
//	type RejectAuthRequest struct {
//		DocumentID     int    `json:"document_id"`
//		UserID         int    `json:"user_id"`
//		TeacherComment string `json:"teacher_comment"` // 教員コメント
//	}

func RejectAuth(c *gin.Context) {

	request := model.RejectAuthRequest{}
	post := model.Post{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print("&")

	//DB接続
	db := infra.DBInitGorm()

	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post)

	if request.TeacherComment != "" && post.PostID == 2 {
		//却下処理
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
	} else if request.TeacherComment == "" {
		//教員コメントがない場合
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "TEACHER COMMENT NOT EXISTS",
		})
	} else {
		//却下者が不適切な場合
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "POST ERROR",
		})
	}

}
