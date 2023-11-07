package api

import (
	"fmt"
	"net/http"

	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

//使用する構造体(model.goより)

// type ReadAuthListRequest struct {
// 	UserID int `json:"user_id"` //ユーザID
// }
// type ReadAuthListResponse struct {
// 	ClassName    string `json:"class_name"`
// 	UserName     string `json:"user_name"`
// 	DivisionName string `json:"division_name"`
// 	DocumentID   int    `json:"document_id"`
// 	Status       int    `json:"status"`
// }
// type TakeClassID struct {
// 	PostID  int    `json:"post_id"`
// 	ClassID string `json:"class_id"`
// }

func ReadAuthList(c *gin.Context) {

	response := []model.ReadAuthListResponse{}
	take_class_id := model.TakeClassID{}

	//引数用・引数がなくなったら消す
	request := model.ReadAuthListRequest{}
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
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&take_class_id)
	// 実行予定のSQL
	// SELECT post_id, class_id FROM user WHERE user_id = 2;
	fmt.Println(take_class_id)

	post := take_class_id.PostID

	if post == 1 {
		//学生であるとき
		//学生自身の提出した公欠届を取得
		db.Table("oa").
			Select(
				"cs.class_name",
				"user.user_name",
				"dv.division_name",
				"oa.document_id",
				"oa.status").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", post).
			Where("oa.user_id = ?", request.UserID).
			Scan(&response)
		if db.Error != nil {
			fmt.Print("STUDENT DATA CATCH ERROR!")
		}

	} else if post == 2 {

		//教員の場合
		db.Table("oa").
			Select(
				"cs.class_name",
				"user.user_name",
				"dv.division_name",
				"oa.document_id",
				"oa.status").
			Joins("JOIN user ON oa.user_id = user.user_id").
			Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
			Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
			Where("oa.status = ?", post-1).
			Scan(&response)
		// 実行予定のSQL
		// SELECT cs.class_name, user.user_name, dv.division_name, oa.document_id FROM oa JOIN user ON oa.user_id = user.user_id JOIN division AS dv ON oa.division_id = dv.division_id JOIN classification AS cs ON user.class_id = cs.class_id WHERE oa.status = 2;
		if db.Error != nil {
			fmt.Print("HIGHER TEACHER DATA CATCH ERROR!")
		}
	} else {
		//不適切な役職の場合(含：管理者)
		c.JSON(http.StatusOK, gin.H{"document": "POST ERROR"})
		return
	}

	// ステータス200と、responseを返します
	c.JSON(http.StatusOK, gin.H{"document": response})
	return

}

//担任教員であるとき
//担任クラスのみ出力する
// db.Table("oa").
// 	Select(
// 		"cs.class_name",
// 		"user.user_name",
// 		"dv.division_name",
// 		"oa.document_id",
// 		"oa.status").
// 	Joins("JOIN user ON oa.user_id = user.user_id").
// 	Joins("JOIN division AS dv ON oa.division_id = dv.division_id").
// 	Joins("JOIN classification AS cs ON user.class_id = cs.class_id").
// 	Where("oa.status = ?", post-1).
// 	Where("cs.class_id = ?", take_class_id.ClassID).
// 	Scan(&response)
// // 実行予定のSQL
// // SELECT cs.class_name, user.user_name, dv.division_name, oa.document_id FROM oa JOIN user ON oa.user_id = user.user_id JOIN division AS dv ON oa.division_id = dv.division_id JOIN classification AS cs ON user.class_id = cs.class_id WHERE oa.status = 1 AND cs.class_id = 117;
// if db.Error != nil {
// 	fmt.Print("CLASS TEACHER DATA CATCH ERROR!")
// }
