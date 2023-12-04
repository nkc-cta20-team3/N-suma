package student

import (
	"errors"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DocumentStatus struct {
	Status int
}

func ReadDocument(c *gin.Context) {

	request := model.ReadDocumentRequest{}
	response := model.ReadDocumentResponse{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Print(request.DocumentID)

	//データベース接続
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	/*
		type ReadDocumentResponse struct {
			DocumentID int `json:"document_id"`
			RequestDate time.Time `json:"request_date"`
			StudentID int `json:"student_id"`
			ClassName string `json:"class_name"`
			StudentName string `json:"student_name"`
			StartDate time.Time `json:"absence_start_date"`
			StartFlame int `json:"start_flame"`
			EndDate time.Time `json:"end_date"`
			EndFlame int `json:"end_flame"`
			Location string `json:"location"`
			StudentComment string `json:"student_comment"`
			TeacherComment string `json:"teacher_comment"`
		}
	*/

	//引数定義
	post := model.Post{}
	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post)

	PostID := post.PostID

	// fmt.Println(PostID)
	//管理者処理
	if PostID == 0 {
		// fmt.Println("権限がありません")
		c.JSON(http.StatusBadRequest, gin.H{"document": "権限がありません"})
		return
	} else if PostID == 1 {
		//学生処理

		// データベースからデータを取得する
		err := db.Debug().Table("oa").
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
				"cs.class_name",
				"user.user_name").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("document_id = ?", request.DocumentID).
			Where("oa.user_id = ?", request.UserID).
			First(&response).Error
		if db.Error != nil {
			errMsg := "データベースからデータを取得できませんでした"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"document": "書類が見つかりませんでした"})
			return
		}

		//書類ステータス取得

		document := DocumentStatus{}
		err = db.Table("oa").Select("status").Where("document_id = ?", request.DocumentID).First(&document).Error

		//認可済みまたは却下状態かどうかを判別
		if document.Status == 2 || document.Status == -1 {
			//既読処理
			read_flag := false
			err = db.Table("oa").Where("document_id = ?", request.DocumentID).Updates(model.UpdateDocument{ReadFlag: &read_flag}).Error
			if err != nil {
				//エラーハンドリング
				c.JSON(http.StatusInternalServerError, gin.H{"document": "READ UPDATE ERROR"})
			}
		}

		// log.Println(response)

		c.JSON(http.StatusOK, gin.H{"document": response})
	} else {

		c.JSON(http.StatusBadRequest, gin.H{"document": "POST ERROR"})
		return
	}
}
