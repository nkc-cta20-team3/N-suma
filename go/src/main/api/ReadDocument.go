package api

import (
	"encoding/json"
	"log"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadDocument(c *gin.Context) {

	/*
	request := model.ReadDocumentRequest{}
	response := model.ReadDocumentResponse{}

	// パラメータ取得
	document_id := c.Param("document_id")
	if document_id == "" {
		errMsg := "document_idが空です"
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	

	//データベース接続
	db := infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}


	RequestDate string `json:"request_date"`
	StudentID int `json:"student_id"`
	ClassName string `json:"class_name"`
	StudentName string `json:"student_name"`
	StartDate string `json:"start_date"`
	StartFlame int `json:"start_flame"`
	EndDate string `json:"end_date"`
	EndFlame int `json:"end_flame"`
	Location string `json:"location"`
	StudentComment string `json:"student_comment"`
	TeacherComment string `json:"teacher_comment"`

	// データベースからデータを取得
	db.Table("absence_document").Select(
		"request_date",
		"student_id",
		"company_id",
		"reason_id",
		"absence_start_date",
		"absence_start_flame",
		"absence_end_date",
		"absence_end_flame",
		"location",
		"read_flag",
		"status",
		"student_comment",
		"teacher_comment").Where("document_id = ?", document_id).First(&document)
	if db.Error != nil {
		errMsg := "データベースからデータを取得できませんでした"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	log.Println(document)

	// JSONに変換
	json, err := json.Marshal(document)
	if err != nil {
		errMsg := "JSONに変換できませんでした"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}
	log.Println(string(json))
	*/

	// JSONで返す
	// c.JSON(http.StatusOK, document)
	c.JSON(http.StatusOK, "OK")
	
}
