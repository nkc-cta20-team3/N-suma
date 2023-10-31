package api

import (
	"main/infra"

	"github.com/gin-gonic/gin"
)

func ReadAlarm(c *gin.Context) {

	//DB接続
	db := infra.DBInitGorm()
}
