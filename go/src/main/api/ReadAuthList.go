package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadAuthList(c *gin.Context) {

	request := model.ReadAuthListRequest{}
	response := []model.ReadAuthListResponse{}
	take_class_id := model.TakeClassId{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//DB接続
	db := infra.DBInitGorm()

	//サブクエリ(副問い合わせ)の作成

	//所属クラス名を取得
	db.Table("user").Select("post_id, class_id").Where("user_uid = ?", request.UserUid).First(&take_class_id)
	post := take_class_id.PostId - 1

	

	if post == 1 {

		//担任教員であるとき
		//担任クラスのみ出力する
		db.Table("oa").
			Select(
				"cs.class_name",
				"user.user_name",
				"dv.division_name",
				"oa.document_id").
			Joins("JOIN user ON oa.user_uid = user.user_uid").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN class AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", 0).
			Where("cs.class_id = ?", take_class_id.ClassId).
			Scan(&response)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}

	} else if post >= 2 && post <= 5 {

		//主任以上の場合
		//すべてのクラスを出力する

		db.Table("absence_document AS ad").
			Select(
				"st.class_name",
				"st.student_name",
				"ar.absence_category",
				"ad.document_id").
			Joins("JOIN students AS st ON ad.student_id = st.student_id").
			Joins("JOIN absence_reason AS ar ON ad.reason_id = ar.reason_id").
			Where("ad.status = ?", post).
			Scan(&response)
		if db.Error != nil {
			fmt.Print("ERROR!")
		}

	} 

	// ステータス200と、responseを返します
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}
