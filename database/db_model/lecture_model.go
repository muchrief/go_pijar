package db_model

type Lecture struct {
	Id           int64     `gorm:"primaryKey"`
	SchoolName   string    `gorm:"notNull" json:"school_name"`
	SupervisorId int64     `gorm:"notNull" json:"supervisor_id"`
	Name         string    `gorm:"notNull" json:"name"`
	Title        string    `gorm:"notNull" json:"title"`
	Supervisor   *Lecture  `gorm:"foreignKey:SupervisorId;references:id;" json:"supervisor,omitempty"`
	School       *School   `json:"school,omitempty"`
	Courses      []*Course `gorm:"many2many:lecture_courses" json:"courses,omitempty"`
}

// CREATE TABLE LECTURER(
// 	STFID INTEGER NOT NULL,
// 	SCHLNM VARCHAR(30) NOT NULL,
// 	SUPID INTEGER,
// 	LECTNM VARCHAR(20),
// 	LECTTITL VARCHAR(30),
// 	OFFROOM VARCHAR(10),
// 	PRIMARY KEY (STFID),
// 	FOREIGN KEY (SCHLNM) REFERENCES SCHOOL (SCHLNM),
// 	FOREIGN KEY (SUPID) REFERENCES LECTURER (STFID)
//    )
