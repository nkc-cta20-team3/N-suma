package api

import (
	"net/http"
	"fmt"
	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

type Post struct {
	PostID int
}


func UpdateUser(c *gin.Context) {

	request := model.UpdateUserRequest{}

	//POSTで受け取った値を格納する
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//DBに接続
    db := infra.DBInitGorm()
	//引数定義
	post := Post{}
	//post_idを取得(reqest.UserIDの部分に取得したいユーザのユーザIDを入れる)
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).First(&post)

	PostID := post.PostID 

	fmt.Println(PostID)

	//管理者処理
    if PostID == 0 {
        
		//更新処理
		db.Table("user").
		Where("user_id = ?", request.UserID).
		Updates(model.UpdateUserRequest{
			UserID: request.UserID,
			UserName: request.UserName,
			UserNumber: request.UserNumber,
			PostID: request.PostID,
			ClassID: request.ClassID,
			MailAddress: request.MailAddress})
	
		//エラーハンドリング
		if db.Error != nil {
			errMsg := "UPDATE ERROR"
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusOK, gin.H{"flag": "データの更新が成功しました"})
	
	//新たに管理者を増やすことができない
	/*

	*/
	} else {
		fmt.Println("権限がありません")
		
	}
		  

}