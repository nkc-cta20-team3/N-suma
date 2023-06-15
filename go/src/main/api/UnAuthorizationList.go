package api

import (
	"fmt"
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type TakeClassName struct {
	ClassName string `json:"class_name"`
}

func UnAuthorizationList(c *gin.Context) {

	// 確認用
	log.Println("UnAuthorizationList")

	//構造体定義

	//POST用
	teacher_data := model.TeacherData{}
	document := []model.UnAuthorizeList{}
	take_class_name := TakeClassName{}

	//POST用、値を格納する
	if err := c.ShouldBindJSON(&teacher_data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 確認用
	log.Println("Post Data")

	//DB接続(共通)
	db := infra.DBInitGorm()

	//サブクエリ(副問い合わせ)の作成

	//所属クラス名を取得
	db.Table("teachers").Select("class_name").Where("teacher_id = ?", teacher_data.TeacherID).First(&take_class_name)

	log.Println(take_class_name)
	position := teacher_data.Position - 1
	log.Println(teacher_data)

	if teacher_data.Position == 1 {
		//担任教員であるとき

		log.Println("aaa")

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", position).
			Where("st.class_name = ?", take_class_name.ClassName).
			Scan(&document)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}
	} else {

		//主任以上の場合
		log.Println("bbb")

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", &teacher_data.Position).
			Scan(&document)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}
	}

	//subquery := db.Table("class").Select("class_name").Where("class_name = ?", class_name)

	//POST用
	// db.Table("absence_document AS ad").
	// 	Select(
	// 		"st.class_name",
	// 		"st.student_name",
	// 		"ar.absence_category",
	// 		"ad.document_id").
	// 	Joins("JOIN students AS st ON ad.student_id = st.student_id").
	// 	Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
	// 	Where("ad.status = ?", &teacher_data.Position).
	// 	Scan(&document)
	// if db.Error != nil {
	// 	fmt.Print("ERROR!")
	// }

	payload := gin.H{
		"document": document,
	}
	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
