package db_model

import "time"

type Student struct {
	Id          string          `gorm:"primaryKey" json:"id"`
	Nim         string          `gorm:"notNull;unique" json:"nim"`
	Name        string          `gorm:"notNull" json:"name"`
	Birthday    time.Time       `gorm:"notNull" json:"birthday"`
	ProgramCode string          `gorm:"notNull" json:"programme_code"`
	YearEnroll  int64           `gorm:"notNull" json:"year_enroll"`
	Courses     []CourseStudent `json:"courses"`
	Programme   *Programme      `json:"programme,omitempty"`
}

// CREATE TABLE STUDENT(
// 	STUID INTEGER NOT NULL,
// 	PROGCD CHAR(11) NOT NULL,
// 	STUNM VARCHAR(30),
// 	STUBRTH DATE,
// 	YRENRL INTEGER,
// 	PRIMARY KEY (STUID),
// 	FOREIGN KEY (PROGCD) REFERENCES PROGRAMME (PROGCD)
//    )
