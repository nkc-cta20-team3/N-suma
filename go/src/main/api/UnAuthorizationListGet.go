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
	text := ""

	if 1 == 1 {
		text = "st.class_name = ?"
	}
	class_name := "CTA20"

	//ここに副問い合わせを入れる
	//subquery := db.Table("")

	//DB接続
	db := infra.DBInitGorm()
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
		Where(text, class_name).
		Scan(&document)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	payload := gin.H{
		"document": document,
	}

	//ステータスと
	c.JSON(http.StatusOK, payload)

}
