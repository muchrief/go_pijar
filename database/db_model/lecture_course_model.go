package db_model

type LectureCourse struct {
	LectureId  int64  `gorm:"primaryKey"`
	CourseCode string `gorm:"primaryKey"`
}

// CREATE TABLE LECTURER_COURSE(
// 	STFID INTEGER NOT NULL,
// 	CRSECD CHAR(10) NOT NULL,
// 	PRIMARY KEY (STFID, CRSECD),
// 	FOREIGN KEY (CRSECD) REFERENCES COURSE (CRSECD),
// 	FOREIGN KEY (STFID) REFERENCES LECTURER (STFID)
//    )
