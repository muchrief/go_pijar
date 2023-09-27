package db_model

type Programme struct {
	Code       string   `gorm:"primaryKey" json:"code"`
	Title      string   `gorm:"notNull" json:"title"`
	Level      string   `gorm:"notNull" json:"level"`
	Duration   string   `gorm:"notNull" json:"duration"`
	SchoolName string   `gorm:"notNull" json:"school_name"`
	School     *School  `json:"school,omitempty"`
	Courses    []Course `gorm:"foreignKey:ProgrammeCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"courses"`
}

// CREATE TABLE PROGRAMME(
// 	PROGCD CHAR(11) NOT NULL,
// 	SCHLNM VARCHAR(30) NOT NULL,
// 	PROGTITL VARCHAR(20),
// 	PROGLVL VARCHAR(10),
// 	PROGLEN VARCHAR(20),
// 	PRIMARY KEY (PROGCD),
// 	FOREIGN KEY (SCHLNM) REFERENCES SCHOOL (SCHLNM)
//    )
