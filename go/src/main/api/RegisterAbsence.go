package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main/infra"
	"main/model"
)

func RegisterAbsence(c *gin.Context) {

	// text := c.Param("absence_list")
	
	//データベース接続
	engine:= infra.DBInit()

	//エラー出力
	if engine == nil {
		errMsg := "データベース接続の初期化に失敗しました"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	//select文test 構造体をimportしてそのリストに全ての情報を入れる
	var users []model.Users
	err := engine.Table("users").Find(&users)
	if err != nil {
		errMsg := "データの取得に失敗しました"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	c.JSON(http.StatusOK, users)

	//データベース切断
	defer engine.Close()
}
