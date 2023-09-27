package db_model

type Campus struct {
	Id          string     `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"notNull;unique" json:"name"`
	Address     string     `gorm:"notNull" json:"address"`
	Distance    int64      `gorm:"notNull" json:"distance"`
	Abbrevation string     `gorm:"notNull;unique" json:"abbrevation"`
	Faculties   []*Faculty `gorm:"foreignKey:CampusId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"faculties,omitempty"`
	Schools     []*School  `gorm:"foreignKey:CampusId;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"schools,omitempty"`
	Club        *Club      `gorm:"foreignKey:CampusName;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"club,omitempty"`
}

// CREATE TABLE CAMPUS(
// 	CMPSNM VARCHAR(20) NOT NULL,
// 	CMPSADDR VARCHAR(50),
// 	DIST NUMERIC(8, 2),
// 	BUSNO VARCHAR(10),
// 	PRIMARY KEY (CMPSNM)
//    )
