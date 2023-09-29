package repo

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
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

func (c *CampusRepo) GetCampus() ([]*db_model.Campus, error) {
	var result []*db_model.Campus
	if err := c.db.Model(&db_model.Campus{}).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CampusRepo) GetCampusDetail(id string) (*db_model.Campus, error) {
	var result db_model.Campus
	if err := c.db.Model(&db_model.Campus{
		Id: id,
	}).
		Preload("Faculties").
		Preload("Schools").
		Preload("Club").
		First(&result).Error; err != nil {
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

func (c *CampusRepo) GetCampusSchools(id string) (*db_model.Campus, error) {
	var result db_model.Campus
	if err := c.db.Model(&db_model.Campus{
		Id: id,
	}).Preload("Schools").First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CampusRepo) Campus(p *model.Pagination) ([]*db_model.Campus, error) {
	var result []*db_model.Campus
	err := c.db.Scopes(
		helper.Paginate(p, c.db),
	).Find(&result).Error

	return result, err
}
