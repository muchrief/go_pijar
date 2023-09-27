package db_model

type Club struct {
	Id         string   `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"notNull" json:"name"`
	Phone      string   `gorm:"notNull" json:"phone"`
	CampusName string   `gorm:"notNull" json:"campus_name"`
	Sports     []*Sport `gorm:"foreignKey:SportName;references:name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"sports"`
	Campus     *Campus  `json:"campus,omitempty"`
}

// CREATE TABLE CLUB(
// 	CLBNM VARCHAR(20) NOT NULL,
// 	CMPSNM VARCHAR(20) NOT NULL,
// 	CLBBLD VARCHAR(20),
// 	PHNNO CHAR(12),
// 	PRIMARY KEY (CLBNM),
// 	FOREIGN KEY (CMPSNM) REFERENCES CAMPUS (CMPSNM)
//    )
