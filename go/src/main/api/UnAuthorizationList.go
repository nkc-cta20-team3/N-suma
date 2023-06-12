package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func UnAuthorizationList(c *gin.Context) {

	//引数を取得する
	status := c.Param("teacher_data")

	//構造体定義
	document := []model.UnAuthorizeList{}
	//DB接続
	db := infra.DBInitGorm()

	db.Table("absence_document AS ad").
		Select(
			"st.class_name",
			"st.student_name",
			"ar.absence_category",
			"ad.document_id").
		Joins("JOIN students AS st ON ad.student_id = st.student_id").
		Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
		Where("ad.status = ?", status).
		Scan(&document)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, document)
	// c.JSON(http.StatusOK, aaa)

}
