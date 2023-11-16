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

	db.Table("user").Select("post_id").Where("user_id = ?", request.UserID).Scan(&post)
}
