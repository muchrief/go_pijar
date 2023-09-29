package repo

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
	"gorm.io/gorm"
)

type ClubRepo struct {
	db *gorm.DB
}

func NewClubRepo(db *gorm.DB) *ClubRepo {
	return &ClubRepo{
		db: db,
	}
}

func (c *ClubRepo) Clubs(p *model.Pagination) ([]*db_model.Club, error) {
	var result []*db_model.Club
	err := c.db.Scopes(
		helper.Paginate(p, c.db),
	).
		// Preload("Campus").
		Preload("Sports").
		Find(&result).Error

	return result, err
}

func (c *ClubRepo) Club(id string) ([]*db_model.Club, error) {
	var result []*db_model.Club
	err := c.db.Where(&db_model.Club{
		Id: id,
	}).
		Preload("Campus").
		Preload("Sport").
		Find(&result).Error

	return result, err
}
