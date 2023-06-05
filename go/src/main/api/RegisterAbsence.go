package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main/infra"
	"main/model"
)

func RegisterAbsence(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	//HTTPリクエストのメソッドがPOSTメゾットあるかチェック
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//データベース接続
	engine:= infra.DBInit()
	if engine == nil {
		errMsg := "データベース接続の初期化に失敗しました"
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}

	//データの受け取り
	var doc model.AbsenceDocument
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 挿入時の時間を取得
	// createdAt := time.Now()
	
	// //構造体にそのまま打ち込む場合
	// newDate := model.AbsenceDocument{
	// 	StudentID:           int
	// 	CompanyID:           int
	// 	ReasonID:            int
	// 	RequestDate:         time.Time
	// 	AbsenceStartDate:    time.Time
	// 	AbsenceStartFlame:   int
	// 	AbsenceEndDate:      time.Time
	// 	AbsenceEndFlame:     int
	// 	Location:            string
	// 	ReadFlag:            bool
	// 	Status:              int
	// 	StudentInputComment: sql.NullString
	// 	TeacherInputComment: sql.NullString
	// }

	_, err = engine.Table("absence document").Insert(&doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データの挿入に失敗しました"})
		return
	}

	//成功ステータス
	c.JSON(http.StatusOK, "成功しました")

	//データベース切断
	defer engine.Close()
}
