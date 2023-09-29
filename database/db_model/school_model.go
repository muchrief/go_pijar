package db_model

type School struct {
	Id          string       `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"notNull;unique" json:"name"`
	Abbrevation string       `gorm:"notNull;unique" json:"abbrevation"`
	CampusId    string       `gorm:"notNull" json:"campus_id"`
	FacultyId   string       `gorm:"notNull" json:"faculty_id"`
	Campus      *Campus      `json:"campus,omitempty"`
	Faculty     *Faculty     `json:"faculty,omitempty"`
	Programmes  []*Programme `gorm:"foreignKey:SchoolName;references:name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"programmes"`
	Lectures    []*Lecture   `gorm:"foreignKey:SchoolName;references:name;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"lectures"`
}

// CREATE TABLE SCHOOL(
// 	SCHLNM VARCHAR(30) NOT NULL,
// 	CMPSNM VARCHAR(20) NOT NULL,
// 	FACNM VARCHAR(30) NOT NULL,
// 	SCHLBLD VARCHAR(20),
// 	PRIMARY KEY (SCHLNM),
// 	FOREIGN KEY (CMPSNM) REFERENCES CAMPUS (CMPSNM),
// 	FOREIGN KEY (FACNM) REFERENCES FACULTY (FACNM)
//    )
