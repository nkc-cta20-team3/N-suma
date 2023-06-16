package api

import (
	"log"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更新用構造体
// type UpdateAbsence struct {
// 	Status         int    `json:"status"`
// 	TeacherComment string `json:"teacher_comment"`
// }

// 更新前表
type AbsenceDocument struct {
	DocumentID        int
	StudentID         int
	CompanyID         int
	ReasonID          int
	RequestDate       string
	AbsenceStartDate  string
	AbsenceStartFlame int
	AbsenceEndDate    string
	AbsenceEndFlame   int
	Location          string
	ReadFlag          bool
	Status            int
	StudentComment    string
	TeacherComment    string
}

func AuthorizeAbsence(c *gin.Context) {

	//引数取得
	document := model.AuthorizeAbsence{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//必要な変数を定義
	var documentStatus int
	var teacherPosition int
	//var updateAbsence UpdateAbsence
	var message string = ""

	//認可ステータスの取得(導通未確認)
	db.Table("absence_document").Select("status").Where("document_id = ?", document.DocumentID).Scan(&documentStatus)

	//役職IDの取得(導通未確認)
	db.Table("teachers").Select("position_id").Where("teacher_id = ?", document.TeacherID).Scan(&teacherPosition)

	log.Println(document)
	log.Println(documentStatus)
	log.Println(teacherPosition)

	if teacherPosition == documentStatus+1 {
		//認可できる状態

		db.Table("absence_document").
			Where("document_id = ?", document.DocumentID).
			Updates(AbsenceDocument{Status: teacherPosition, TeacherComment: document.TeacherComment})
			// 	Update("status", documentStatus).
			// 	Update("teacher_comment", document.TeacherComment)

			// db.Model(&AbsenceDocument{}).
			// 	Where("document_id = ?", document.DocumentID).
			// 	Update("status", documentStatus).
			// 	Update("teacher_comment", document.TeacherComment)

		// db.Table("absence_document").Select("status,teacher_comment").Where("document_id = ?", document.DocumentID).First(&updateAbsence)
		// updateAbsence.Status = documentStatus
		// updateAbsence.TeacherComment = document.TeacherComment

		//db.Save(&updateAbsence)
		//err := db.Save(&updateAbsence)
		// if err.Error != nil {
		// 	c.JSON(400, gin.H{"error": err.Error()})
		// 	return
		// }

		message = "SUCCESS!"

	}

	//JSONに変えるときに使う
	payload := gin.H{
		"message": message,
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
