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

	var user *db_model.User

	err := us.db.Transaction(func(tx *gorm.DB) error {
		userRepo := repo.NewUserRepo(tx)
		userDb, err := userRepo.GetUserByEmail(email)
		if err != nil {
			if err == repo.ErrUserNotFound {
				user, err = userRepo.CreateUser(payload)
				return err
			}
			return err
		}
		user = userDb

		return nil
	})

	return user, err
}
