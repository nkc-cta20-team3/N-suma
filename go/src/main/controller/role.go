package contoroller

import (
	"log"
	"main/infra"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	PostName string
}

func RoleMiddleware(role string, UUID string) gin.HandlerFunc {
	return func(c *gin.Context) {

		//受け取り用構造体
		post := Post{}

		//取得するユーザのUUID
		request_id = UUID

		//DB接続
		db := infra.DBInitGorm()

		err := db.Table("user").
			Select("post.post_name").
			Joins("JOIN post ON user.post_id = post.post_id").
			Where("user_uuid = ?", request_id).Scan(&post).Error
		if err != nil {
			//エラーハンドリング
			//役職が登録されていない場合
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}

		//引数と役職が一致するか判定
		if role == Post.PostName {
			//引数と役職が一致
			log.Printf("ROLL NAME : %s\n", post.PostName)
			c.Next()
		} else {
			//役職が一致しない
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error": "ROLE AUTHORISE ERROR",
			})
			return
		}
	}
}
