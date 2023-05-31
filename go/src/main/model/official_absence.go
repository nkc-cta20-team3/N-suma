package main

import (
	"database/sql"
	"time"
)

type AbsenceDocument struct {
	DocumentID           int            `xorm:"documentid pk autoincr"`
	StudentID            int            `xorm:"studentid fk"`
	CompanyID            int            `xorm:"companyid fk"`
	ReasonID             int            `xorm:"reasonid  fk"`
	RequestDate          time.Time      `xorm:"requestdate"`
	AbsenceStartDate     time.Time      `xorm:"absencestartdate"`
	AbsenceStartFlame    int            `xorm:"absencestartflame"`
	AbsenceEndDate       time.Time      `absenceenddate"`
	AbsenceEndFlame      int            `xorm:"absenceendflame"`
	Location             string         `xorm:"location"`
	ReadFlag             bool           `xorm:"readflag"`
	Status               int            `xorm:"status"`
	StudentInputComment  sql.NullString `xorm:"studentinputcomment"`
	TeacherInputComment  sql.NullString `xorm:"teacherinputcomment"`
}

