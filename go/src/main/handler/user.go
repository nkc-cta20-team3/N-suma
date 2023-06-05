package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"main/model"
	"main/service"
)

func Create(c *gin.Context) {
	input := model.UserInput{}

	// ヘッダーのJSONをinputにバインドします
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// サービス層をNewします
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	// サービス層のCreateメソッドを呼び出します
	payload, err := user.Create(&input)
	if err != nil {
		log.Fatal(err)
	}

	// ステータス200と、payloadを返します
	c.JSON(http.StatusOK, payload)
}

func GetAll(c *gin.Context) {
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetOne(c *gin.Context) {
	// user-idのパラメータを数字に変換します
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetOne(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context) {
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	input := model.UserInput{}
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	payload, err := user.Update(&input, userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Delete(c *gin.Context) {
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	err = user.Delete(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "削除されました")
}