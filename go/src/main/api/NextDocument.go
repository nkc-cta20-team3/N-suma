package api

import (
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NextDocument(c *gin.Context) {
	//必要な変数(引数・戻り値)の定義
	request := []model.NextDocumentRequest{}
	response := []model.NextDocumentResponse{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DBに接続
	db := infra.DBInitGorm()

}
