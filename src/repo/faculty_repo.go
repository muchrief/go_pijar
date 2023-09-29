package repo

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
	"gorm.io/gorm"
)

type FacultyRepo struct {
	db *gorm.DB
}

func NewFacultyRepo(db *gorm.DB) *FacultyRepo {
	return &FacultyRepo{
		db: db,
	}
}

func (f *FacultyRepo) Faculties(p *model.Pagination) ([]*db_model.Faculty, error) {
	var result []*db_model.Faculty
	err := f.db.Model(&db_model.Faculty{}).
		Scopes(
			helper.Paginate(p, f.db),
		).
		// Preload("Campus").
		Preload("Committe").
		Find(&result).Error

	return result, err
}

func (f *FacultyRepo) FacultySchools(id string) ([]*db_model.Faculty, error) {
	var result []*db_model.Faculty
	err := f.db.Model(&db_model.Faculty{}).
		Where(&db_model.Faculty{
			Id: id,
		}).
		Preload("Schools").
		Find(&result).Error

	return result, err
}
