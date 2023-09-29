package db_model

type Faculty struct {
	Id              string           `gorm:"primaryKey" json:"id"`
	Name            string           `gorm:"notNull;unique" json:"name"`
	DeanName        string           `gorm:"notNull" json:"dean_name"`
	Abbrevation     string           `gorm:"notNull;unique" json:"abbrevation"`
	CampusId        string           `json:"campus_id"`
	Schools         []*School        `gorm:"foreignKey:FacultyId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"schools,omitempty"`
	Committe        *Committe        `gorm:"foreignKey:FacultyName;references:name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"committe,omitempty"`
	CommitteLecture *CommitteLecture `gorm:"foreignKey:FacultyName;references:name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"committe_lecture,omitempty"`
	// Campus          *Campus          `json:"campus,omitempty"`
}

// CREATE TABLE FACULTY(
// 	FACNM VARCHAR(30) NOT NULL,
// 	DEANNM VARCHAR(20),
// 	FACBLD VARCHAR(20),
// 	PRIMARY KEY (FACNM)
//    )
