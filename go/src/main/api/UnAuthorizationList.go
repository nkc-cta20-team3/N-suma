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

func UnAuthorizationList(c *gin.Context) {

	document_id := c.Param("teacher_data")
	// input := model.UserInput{}

	var document = Document2{}
	db := infra.DBInitGorm()

	db.Table("absence_document").Select(
		"document_id",
		"student_id",
		"location").Where("document_id > ?", document_id).First(&document)
	if db.Error != nil {
		fmt.Print("ERROR!")
	}

	fmt.Printf("document_id: %d, student_id: %d, location: %s\n", document.DocumentID, document.StudentID, document.Location)

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
