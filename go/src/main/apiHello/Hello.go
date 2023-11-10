package apiHello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GETメソッドで動作し、ステータス200と、HelloWorldを返します
func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "HelloWorld"})
	return

}

// /hello/{name}
//GETメソッドで動作し、ステータス200と、Helloと取得したパスパラメータを結合した文字列を返します
func Name(c *gin.Context) {

	// パスパラメータを取得します
	name := c.Param("name")

	// ステータス200と、Helloと取得したパスパラメータを結合した文字列を返します
	c.JSON(http.StatusOK, gin.H{"message": "Hello" + name})
	return

}

// GETメソッドで動作し、ステータス200と、現在時刻を返します
func Time(c *gin.Context) {

	// 現在時刻を取得します
	time := time.Now()

	// ステータス200と、現在時刻を返します
	c.JSON(http.StatusOK, gin.H{"message": time})
	return

}

// POSTメソッドで動作し、ステータス200と、POSTされたaとbを合算した値を返します
func Sum(c *gin.Context) {

	// 必要な変数定義
	request := SumRequest{}

	// 引数受け取り
	if err := c.ShouldBindJSON(&request); err != nil {
		// エラーな場合、ステータス400と、エラー情報を返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ステータス200と、POSTされたaとbを合算した値を返します
	c.JSON(http.StatusOK, gin.H{"message": request.A + request.B})
	return

}

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
