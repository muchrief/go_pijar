package db_model

type CourseStudent struct {
	CourseCode    string   `gorm:"primaryKey" json:"course_code"`
	StudentId     string   `gorm:"primaryKey" json:"student_id"`
	YearTaken     int64    `gorm:"notNull" json:"year_taken"`
	SemesterTaken int64    `gorm:"notNull" json:"semester_taken"`
	Grade         int64    `gorm:"notNull" json:"grade"`
	Student       *Student `json:"student,omitempty"`
	Course        *Course  `json:"cource,omitempty"`
}

// CREATE TABLE COURSE_STUDENT(
// 	CRSECD CHAR(10) NOT NULL,
// 	STUID INTEGER NOT NULL,
// 	YRTKN CHAR(4),
// 	SEMTKN VARCHAR(20),
// 	GRDAWRD VARCHAR(10),
// 	PRIMARY KEY (CRSECD, STUID),
// 	FOREIGN KEY (STUID) REFERENCES STUDENT (STUID),
// 	FOREIGN KEY (CRSECD) REFERENCES COURSE (CRSECD)
//    )
