package model

import (
	"database/sql"
	"time"
)

type AbsenceDocument struct {
	StudentID           int            `xorm:"studentid fk"`
	AbsenceStartDate    time.Time      `xorm:"absencestartdate"`
	AbsenceStartFlame   int            `xorm:"absencestartflame"`
	AbsenceEndDate      time.Time      `absenceenddate"`
	AbsenceEndFlame     int            `xorm:"absenceendflame"`
	Location            string         `xorm:"location"`
	StudentInputComment sql.NullString `xorm:"studentinputcomment"`
}
