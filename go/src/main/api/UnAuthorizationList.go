package api

import (
	"net/http"

	// "main/model/???"
	// "main/infra"

	"github.com/gin-gonic/gin"
)

type Teacher_data struct {
	teacher_id int    `json:"teacher_id"`
	position   int    `json:"position"`
	class_name string `json:"class_name"`
}

func UnAuthorizationList(c *gin.Context) {

	aaa := c.Param("teacher_data")

	// input := model.UserInput{}

	// ヘッダーのJSONをinputにバインドします
	// err := c.ShouldBindWith(&input, binding.JSON)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Bad request")
	// 	return
	// }

	// var data Teacher_data

	// err := json.Unmarshal("teacher_data", &data)

	// if err != nil {
	// 	//エラー
	// }

	// fmt.Println(data.teacher_id)
	// fmt.Println(data.position)
	// fmt.Println(data.class_name)

	// JSONに変えるときに使う
	payload := gin.H{
		"message": aaa,
	}

	//DB接続
	// db = infra.DBInit()

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)
	// c.JSON(http.StatusOK, aaa)

}
