package repo

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"gorm.io/gorm"
)

type CampusRepo struct {
	db *gorm.DB
}

func NewCampusRepo(db *gorm.DB) *CampusRepo {
	return &CampusRepo{
		db: db,
	}
}

func (c *CampusRepo) GetCampus(id string) (*db_model.Campus, error) {
	var result db_model.Campus
	if err := c.db.Model(&db_model.Campus{
		Id: id,
	}).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CampusRepo) GetCampusFaculties(id string) (*db_model.Campus, error) {
	var result db_model.Campus
	if err := c.db.Model(&db_model.Campus{
		Id: id,
	}).Preload("Facilties").First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
