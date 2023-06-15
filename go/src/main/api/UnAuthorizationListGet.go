package api

import (
	"fmt"
	"net/http"
	"strconv"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UnAuthorizationListGet(c *gin.Context) {

	//構造体定義
	document := []model.UnAuthorizeList{}

	//引数を取得する
	status, err := strconv.Atoi(c.Param("teacher_data"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//担任の有無によって処理を変える準備
	class_name := "CTA20"

	//DB接続
	db := infra.DBInitGorm()

	//ここに副問い合わせを入れる
	subquery := db.Table("class").Select("class_name").Where("class_name = ?", class_name)
	//SQLの発行
	db.Table("absence_document AS ad").
		Select(
			"st.class_name",
			"st.student_name",
			"ar.absence_category",
			"ad.document_id").
		Joins("JOIN students AS st ON ad.student_id = st.student_id").
		Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
		Where("ad.status = ?", status-1).
		Where("st.class_name IN (?)", subquery).
		Scan(&document)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	payload := gin.H{
		"document": document,
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)

}
