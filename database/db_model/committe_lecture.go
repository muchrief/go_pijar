package db_model

type CommitteLecture struct {
	LectureId     int64    `gorm:"primaryKey" json:"lecture_id"`
	CommitteTitle string   `gorm:"primaryKey" json:"committe_title"`
	FacultyName   string   `gorm:"faculty_name" json:"faculty_name"`
	Faculty       *Faculty `json:"faculty,omitempty"`
	Lecture       *Lecture `json:"lecture"`
}

// CREATE TABLE COMMITTEE_LECTURER(
// 	STFID INTEGER NOT NULL,
// 	COMMTITL VARCHAR(30) NOT NULL,
// 	FACNM VARCHAR(30) NOT NULL,
// 	PRIMARY KEY (STFID, COMMTITL, FACNM),
// 	FOREIGN KEY (STFID) REFERENCES LECTURER (STFID),
// 	FOREIGN KEY (COMMTITL, FACNM) REFERENCES COMMITTEE (COMMTITL,
// 	FACNM)
//    )
