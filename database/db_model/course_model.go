package db_model

type Course struct {
	Code          string           `gorm:"primaryKey" json:"code"`
	ProgrammeCode string           `gorm:"notNull" json:"programme_code"`
	Title         string           `gorm:"notNull" json:"title"`
	Students      []*CourseStudent `gorm:"foreignKey:CourseCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"students,omitempty"`
	Programme     *Programme       `json:"programme,omitempty"`
	Precourse     []*Precourse     `gorm:"foreignKey:CourseCode;references:code;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"precourse,omitempty"`
	Lectures      []*Lecture       `gorm:"many2many:lecture_courses" json:"lectures,omitempty"`
}

// CREATE TABLE COURSE(
// 	CRSECD CHAR(10) NOT NULL,
// 	PROGCD CHAR(11) NOT NULL,
// 	CRSETITL VARCHAR(20),
// 	PRIMARY KEY (CRSECD),
// 	FOREIGN KEY (PROGCD) REFERENCES PROGRAMME (PROGCD)
//    )
