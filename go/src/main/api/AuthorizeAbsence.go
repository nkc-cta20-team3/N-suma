package api

import (
	"log"
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	var message string = "ERROR"

	//認可ステータスの取得
	db.Table("absence_document").Select("status").Where("document_id = ?", document.DocumentID).Scan(&documentStatus)

	//役職IDの取得
	db.Table("teachers").Select("position_id").Where("teacher_id = ?", document.TeacherID).Scan(&teacherPosition)

	log.Println(document)
	log.Println(documentStatus)
	log.Println(teacherPosition)

	if teacherPosition == documentStatus+1 {
		//認可できる状態

		db.Table("absence_document").
			Where("document_id = ?", document.DocumentID).
			Updates(AbsenceDocument{Status: teacherPosition, TeacherComment: document.TeacherComment})

		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		message = "SUCCESS!"

	}

	//JSONに変えるときに使う
	payload := gin.H{
		"message": message,
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
