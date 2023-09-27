package service

import (
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/repo"
	"gorm.io/gorm"
)

type UserService struct {
	db   *gorm.DB
	repo *repo.UserRepo
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db:   db,
		repo: repo.NewUserRepo(db),
	}
}

func (us *UserService) CreateUser(email, password string) (*db_model.User, error) {
	payload := &db_model.User{
		Username: email,
		Password: password,
	}

	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != (&db_model.User{}) {
		return user, nil
	}

	user, err = us.repo.CreateUser(payload)
	return user, err
}
