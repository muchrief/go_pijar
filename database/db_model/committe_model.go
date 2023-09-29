package db_model

type Committe struct {
	Title         string `gorm:"primaryKey"`
	FacultyName   string `gorm:"primaryKey" json:"faculty_name"`
	MeetFrequency int64  `gorm:"notNull;default:3;" json:"meet_freq"`
	// Faculty       *Faculty `json:"faculty,omitempty"`
}

// CREATE TABLE COMMITTEE(
// 	COMMTITL VARCHAR(30) NOT NULL,
// 	FACNM VARCHAR(30) NOT NULL,
// 	MTFREQ VARCHAR(10),
// 	PRIMARY KEY (COMMTITL, FACNM),
//    Master Thesis – Weiguang Zhang McMaster University- Computing and Software
// 	14
// 	FOREIGN KEY (FACNM) REFERENCES FACULTY (FACNM)
//    )
