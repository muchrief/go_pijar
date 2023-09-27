package db_model

type Precourse struct {
	CourseCode    string  `gorm:"primaryKey" json:"course_code"`
	PrecourseCode string  `gorm:"primaryKey" json:"precourse_code"`
	Precourse     *Course `json:"precourse,omitempty"`
}

// CREATE TABLE PRE_COURSE(
// 	CRSECD CHAR(10) NOT NULL,
// 	PRECRSECD CHAR(10) NOT NULL,
// 	PRIMARY KEY (CRSECD, PRECRSECD),
// 	FOREIGN KEY (CRSECD) REFERENCES COURSE (CRSECD),
// 	FOREIGN KEY (PRECRSECD) REFERENCES COURSE (CRSECD)
//    )
