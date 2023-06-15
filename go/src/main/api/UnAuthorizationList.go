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

	// 確認用
	log.Println("UnAuthorizationList")

	//構造体定義

	//使用する構造体
	teacher_data := model.TeacherData{}
	document := []model.UnAuthorizeList{}
	take_class_name := model.TakeClassName{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&teacher_data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 確認用
	log.Println("Post Data")

	//DB接続
	db := infra.DBInitGorm()

	//サブクエリ(副問い合わせ)の作成

	//所属クラス名を取得
	db.Table("teachers").Select("position_id, class_name").Where("teacher_id = ?", teacher_data.TeacherID).First(&take_class_name)
	position := take_class_name.PositionID - 1

	if take_class_name.PositionID == 1 {
		//担任教員であるとき
		//担任クラスのみ出力する

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
	} else if take_class_name.PositionID >= 2 && take_class_name.PositionID <= 5 {

		//主任以上の場合
		//すべてのクラスを出力する

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", position).
			Scan(&document)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}
	} else {
		//例外のとき
		//なにもしない
	}

	payload := gin.H{
		"document": document,
	}
	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
