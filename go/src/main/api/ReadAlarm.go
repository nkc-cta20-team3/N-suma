package api

import (
	"main/infra"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	//必要な変数定義
	request := model.ReadAlarmRequest{}
	var AlarmFlag = false

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//役職を識別

	//ロールごとの処理分け

	//担任→未認可リストの存在有無

	//主任以上→未認可リストの存在有無

	//学生→再提出or(認可完了and未読)

	//結果を返却
	c.JSON(http.StatusOK, gin.H{"document": AlarmFlag})
}
