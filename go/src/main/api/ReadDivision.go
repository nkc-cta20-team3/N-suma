package api

import (
	"main/infra"
	"main/model"

	"github.com/gin-gonic/gin"
)

func ReadDivision(c *gin.Context) {

	request := model.ReadDivisionReqest{}
	post := model.Post{}
	//DB接続
	db := infra.DBInitGorm()

	//役職ID取得
	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)

	if post.PostID == 1 {
		//学生のみアクセス可能

	}
}
