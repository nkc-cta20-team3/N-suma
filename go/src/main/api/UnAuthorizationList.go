package api

import (
	"fmt"
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UnAuthorizationList(c *gin.Context) {

	//構造体定義

	//POST用
	teacher_data := []model.TeacherData{}
	document := []model.UnAuthorizeList{}

	//引数を取得する(GET)
	// status, err := strconv.Atoi(c.Param("teacher_data"))
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	//POST用
	if err := c.ShouldBindJSON(&teacher_data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teacher_data[0].Position -= 1
	log.Println(teacher_data)

	//DB接続(共通)
	db := infra.DBInitGorm()

	//POST用
	db.Table("absence_document AS ad").
		Select(
			"st.class_name",
			"st.student_name",
			"ar.absence_category",
			"ad.document_id").
		Joins("JOIN students AS st ON ad.student_id = st.student_id").
		Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
		Where("ad.status = ?", &teacher_data[0].Position).
		Scan(&document)
	// db.Table("absence_document AS ad").
	// 	Select(
	// 		"st.class_name",
	// 		"st.student_name",
	// 		"ar.absence_category",
	// 		"ad.document_id").
	// 	Joins("JOIN students AS st ON ad.student_id = st.student_id").
	// 	Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
	// 	Where("ad.status = ?", status-1).
	// 	Scan(&document)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, document)
	// c.JSON(http.StatusOK, aaa)

}
