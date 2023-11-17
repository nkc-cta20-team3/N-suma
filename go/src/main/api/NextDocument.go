package api

import (
	"errors"
	"fmt"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// model.go記載の構造体
// type NextDocumentRequest struct {
// 	DocumentID int  `json:"document_id"` //書類ID(基準点)
// 	UserID     int  `json:"user_id"`     //ユーザID(誰の書類？)
// 	NextFlag   bool `json:"next_flag"`   //フラグ
// }
// type NextDocumentResponse struct {
// 	DocumentID     int       `json:"document_id"`     //書類ID
// 	RequestAt      time.Time `json:"request_at"`      //申請日
// 	StartTime      time.Time `json:"start_time"`      //公欠開始日時
// 	StartFlame     int       `json:"start_flame"`     //公欠開始時限
// 	EndTime        time.Time `json:"end_time"`        //公欠終了日時
// 	EndFlame       int       `json:"end_flame"`       //公欠終了時限
// 	Location       string    `json:"location"`        //場所
// 	StudentComment string    `json:"student_comment"` //学生コメント
// 	TeacherComment string    `json:"teacher_comment"` //教員コメント
// 	UserNumber     int       `json:"user_number"`     //学籍番号
// 	ClassAddr      string    `json:"class_addr"`      //クラス略称
// 	DivisionName   string    `json:"division_name"`   //区分名
// }

type DocumentArray struct {
	DocumentID int `json:"document_id"`
}

func NextDocument(c *gin.Context) {
	//必要な変数(引数・戻り値)の定義
	request := model.NextDocumentRequest{}
	response := []model.NextDocumentResponse{}
	documentArray := []DocumentArray{}

	// var testUserID int
	post := model.Post{}
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
	err := db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 行が見つからなかった場合の処理
			c.JSON(http.StatusBadRequest, gin.H{"message": "DOCUMENT NOT FOUND"})
			return
		} else {
			//その他のエラーハンドリング
			c.JSON(http.StatusBadRequest, gin.H{"message": "OTHER ERROR"})
			return
		}
	}
	PostID = post.PostID

	if PostID == 1 {
		pre_id := 0
		var search_id int

		//書類ID一覧を取得
		err = db.Table("oa").Select("document_id").Where("user_id = ?", request.UserID).Scan(&documentArray).Error

		for _, id := range documentArray {

			if request.DocumentID == id.DocumentID && !request.NextFlag {
				//前の書類IDを取得
				search_id = pre_id
			} else if request.DocumentID == id.DocumentID && request.NextFlag {
				//次の書類IDを取得(下記elseifで検索IDを取得させる)
				pre_id = -1
			} else if pre_id == -1 {
				search_id = id.DocumentID
			} else {
				pre_id = id.DocumentID
			}
		}

		fmt.Println("search_id:", search_id)
		fmt.Println(documentArray)

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
				"cs.class_abbr",
				"dv.division_name").
			Joins("JOIN user on oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("document_id = ?", search_id).
			First(&response)
		//戻り値
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})

	} else if PostID == 2 {
		//教員処理
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
				"cs.class_abbr",
				"dv.division_name").
			Joins("JOIN user on oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", PostID-1).
			Order("oa.request_at ASC").
			First(&response)
		if db.Error != nil {
			fmt.Print("SQL ERROR!")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "SQL ERROR",
			})
		}
		//戻り値
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})

	} else {
		//教員でも学生でもなかったとき(不正なデータ)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "POST ERROR",
		})
	}
}

//担任用の処理置き場
// //副問合せ用クエリ
// subquery := db.Table("user").
// 	Select("user.class_id").
// 	Joins("JOIN oa ON oa.user_id = user.user_id").
// 	Where("document_id = ?", request.DocumentID)

// //ドキュメントIDの学生を担当する教員のユーザIDを取得
// db.Table("user").Select("user_id,post_id").Where("class_id = ?", subquery).First(&testUserID)

// //担任かどうかをチェック
// if testUserID == request.UserID {

// 	//認可すべき次の書類を取得・格納する(最新の書類)
// 	db.Table("oa").
// 		Select(
// 			"oa.document_id",
// 			"oa.request_at",
// 			"oa.start_time",
// 			"oa.start_flame",
// 			"oa.end_time",
// 			"oa.end_flame",
// 			"oa.location",
// 			"oa.student_comment",
// 			"oa.teacher_comment",
// 			"user.user_number",
// 			"cs.class_addr",
// 			"dv.division_name").
// 		Joins("JOIN user on oa.user_id = user.user_id").
// 		Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
// 		Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
// 		Where("user.user_id = ?", request.UserID).
// 		Where("oa.status =", PostID-1).
// 		Order("oa.request_at ASC").
// 		First(&response)

// 	//エラーハンドリング
// 	if db.Error != nil {
// 		fmt.Print("SQL ERROR!")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "SQL ERROR",
// 		})
// 	}

// 	//戻り値
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": response,
// 	})

// } else {
// 	//リクエストに不正及び誤りがあるとき(教員情報不一致)
// 	c.JSON(http.StatusBadRequest, gin.H{"document": response})
// }
