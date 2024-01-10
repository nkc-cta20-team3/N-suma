package controller

import (
	"net/http"
	
	"main/model"

	"github.com/gin-gonic/gin"
)

func RoleCheckMiddleware(role int) gin.HandlerFunc {
	return func(c *gin.Context) {

		//引数と役職が一致するか判定
		if role == c.MustGet("PostID").(int) {
			//引数と役職が一致
			c.Next()
		} else {
			//役職が一致しない
			errResponse := model.MessageError{}
			errResponse.Message = "ROLE AUTHORISE ERROR"
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResponse)
			return
		}
	}
}
