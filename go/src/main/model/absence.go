package model

import (
	"database/sql"
	"time"
)

type AbsenceDocument struct {
	DocumentID 			int				`json:"document_id"` 			// ドキュメントID
	StudentID           int				`json:"student_id"` 			// 学籍番号
	CompanyID           int				`json:"company_id"`				// 会社ID
	ReasonID            int				`json:"reason_id"` 				// 理由ID
	RequestDate         time.Time		`json:"request_date"`			// 申請日
	AbsenceStartDate    time.Time		`json:"absence_start_date"` 	// 欠席開始日
	AbsenceStartFlame   int				`json:"absence_start_flame"` 	// 欠席開始時限
	AbsenceEndDate      time.Time		`json:"absence_end_date"` 		// 欠席終了日
	AbsenceEndFlame     int				`json:"absence_end_flame"`		// 欠席終了時限
	Location            string			`json:"location"` 				// 場所
	ReadFlag            bool			`json:"read_flag"` 				// 既読フラグ
	Status              int				`json:"status"` 				// ステータス
	StudentInputComment sql.NullByte	`json:"student_input_comment"` 	// 学生入力コメント
	TeacherInputComment sql.NullByte	`json:"teacher_input_comment"` 	// 教員入力コメント
}
