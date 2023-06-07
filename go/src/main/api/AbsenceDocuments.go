package api

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"main/infra"
	"main/model"
)

func GetAbsenceDocuments(c *gin.Context) {
	
	// パラメータ取得
	document_id := c.Param("document_id")
	if document_id == "" {
		errMsg := "document_idが空です"
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	//データベース接続
	db:= infra.DBInitGorm()
	if db.Error != nil {
		errMsg := "データベース接続エラー"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	// データベースからデータを取得
	document := model.AbsenceDocument{}
	db.Table("absence_document").Select(
		"document_id",
		"student_id",
		"company_id",
		"reason_id",
		"request_date",
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

	// JSONで返す
	c.JSON(http.StatusOK, document)
	
}
