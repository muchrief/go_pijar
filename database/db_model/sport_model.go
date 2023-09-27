package db_model

type Sport struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"notNull" json:"name"`
	ClubName string `gorm:"notNull" json:"club_name"`
	Club     *Club  `json:"club,omitempty"`
}

// CREATE TABLE SPORT(
// 	SPRTNM VARCHAR(20) NOT NULL,
// 	CLBNM VARCHAR(20) NOT NULL,
// 	PRIMARY KEY (SPRTNM, CLBNM),
// 	FOREIGN KEY (CLBNM) REFERENCES CLUB (CLBNM)
//    )
