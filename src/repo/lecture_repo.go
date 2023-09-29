package repo

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
	"gorm.io/gorm"
)

type LectureRepo struct {
	db *gorm.DB
}

func NewLectureRepo(db *gorm.DB) *LectureRepo {
	return &LectureRepo{
		db: db,
	}
}

func (l *LectureRepo) GetLectures(p *model.Pagination) ([]*db_model.Lecture, error) {
	var result []*db_model.Lecture

	err := l.db.Model(&db_model.Lecture{}).Scopes(
		helper.Paginate(p, l.db),
	).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *LectureRepo) LectureDetail(id int64) (*db_model.Lecture, error) {
	var result db_model.Lecture

	err := l.db.
		Model(&db_model.Lecture{}).
		Where(&db_model.Lecture{
			Id: id,
		}).
		Preload("Courses").
		// Preload("School").
		Preload("Supervisor").
		First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *LectureRepo) GetLectureCourse(id int64) (*db_model.Lecture, error) {
	var result db_model.Lecture

	err := l.db.Model(&db_model.Lecture{}).Where(&db_model.Lecture{
		Id: id,
	}).Preload("Courses").First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
