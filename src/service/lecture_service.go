package service

import (
	"sync"

	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"
)

type LectureService struct {
	sync.Mutex
	db *gorm.DB
}

func NewLectureService(db *gorm.DB) *LectureService {
	return &LectureService{
		db: db,
	}
}

func (l *LectureService) AddLecture(payload *db_model.Lecture) (*db_model.Lecture, error) {

	var result *db_model.Lecture
	err := l.db.Transaction(func(tx *gorm.DB) error {
		lectureRepo := repo.NewLectureRepo(tx)
		lecture, err := lectureRepo.CreateLecture(payload)
		if err != nil {
			return err
		}
		result = lecture
		return nil
	})

	return result, err
}

func (l *LectureService) UpdateLecture(payload *db_model.Lecture) (*db_model.Lecture, error) {
	l.Lock()
	defer l.Unlock()

	var result *db_model.Lecture
	err := l.db.Transaction(func(tx *gorm.DB) error {
		lectureRepo := repo.NewLectureRepo(tx)
		lecture, err := lectureRepo.UpdateLecture(payload)
		if err != nil {
			return err
		}
		result = lecture
		return nil
	})

	return result, err

}
