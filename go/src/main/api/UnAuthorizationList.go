package api

import (
	"fmt"
	"net/http"

	"main/infra"

	"github.com/gin-gonic/gin"
)

type Document2 struct {
	DocumentID int    `json:"document_id"` // ドキュメントID
	StudentID  int    `json:"student_id"`  // 学籍番号
	Location   string `json:"location"`    // 場所
}

// 引数の構造体
type TeacherData struct {
	TeacherID int    `json:"teacher_id"`
	Position  int    `json:"position"`
	ClassName string `json:"class_name"`
}

// 戻り値の構造体
type UnAuthorizeList struct {
	ClassName       string `json:"class_name"`       //クラス名称
	StudentName     string `json:"student_name"`     //学生氏名
	AbsenceCategory string `json:"absence_category"` //種別
	DocumentID      int    `json:"document_id"`      //書類ID
}

func UnAuthorizationList(c *gin.Context) {

	//引数を取得する
	status := c.Param("teacher_data")

	// input := model.UserInput{}

	document := []UnAuthorizeList{}
	//document := UnAuthorizeList{}
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

	//確認用出力
	/*
		fmt.Printf("class_name: %s, student_name: %s, absence_category: %s,document_id: %d\n",
			document.ClassName, document.StudentName, document.AbsenceCategory, document.DocumentID)
	*/

	// ヘッダーのJSONをinputにバインドします
	// err := c.ShouldBindWith(&input, binding.JSON)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }

	// var data Teacher_data

	// err := json.Unmarshal("teacher_data", &data)

	// if err != nil {
	// 	//エラー
	// }

	// JSONに変えるときに使う
	/*
		payload := gin.H{
			"message": document_id,
		}*/

	//DB接続
	// db = infra.DBInit()

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, document)
	// c.JSON(http.StatusOK, aaa)

}
