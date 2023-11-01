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
	var take_post_id int
	var AlarmFlag = false
	var result int

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//役職を識別
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&take_post_id)

	//ロールごとの処理分け
	if take_post_id == 1 {
		//学生→再提出or(認可完了and未読)

		db.Table("oa").
			Select("document_id").
			Where("status IN (?)", []int{1, 6}).
			Where("user_id = ?", request.UserID).
			First(&result)

	} else if take_post_id >= 2 && take_post_id <= 6 {
		//教員→未認可リストの存在有無
		db.Table("oa").
			Select("oa.document_id").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", take_post_id-1).
			First(&result)

	} else {
		//エラー
		c.JSON(http.StatusOK, gin.H{"document": "POST ERROR"})
		return
	}

	//結果を返却
	c.JSON(http.StatusOK, gin.H{"document": AlarmFlag})
}
