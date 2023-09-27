package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/muchrief/go_pijar/database/db_model"
	"github.com/muchrief/go_pijar/src/helper"
	"github.com/muchrief/go_pijar/src/model"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(payload *db_model.User) (*db_model.User, error) {
	if err := u.db.Create(payload).Error; err != nil {
		return nil, err
	}

	return payload, nil
}

func (u *UserRepo) GetAllUser() ([]*db_model.User, error) {
	var user []*db_model.User
	res := u.db.Find(&user)
	if res.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, res.Error) {
			return nil, ErrUserNotFound
		}

		return nil, res.Error
	}

	return user, nil
}

func (u *UserRepo) GetUser(id string) (*db_model.User, error) {
	var user db_model.User
	res := u.db.Model(&db_model.User{}).Where("id = ?", id).First(&user)
	if res.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, res.Error) {
			return nil, ErrUserNotFound
		}

		return nil, res.Error
	}

	return &user, nil
}

func (u *UserRepo) GetUserByEmail(email string) (*db_model.User, error) {
	var user db_model.User
	res := u.db.Model(&db_model.User{}).Where("username = ?", email).First(&user)
	if res.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, res.Error) {
			return nil, ErrUserNotFound
		}

		return nil, res.Error
	}

	return &user, nil
}

func (u *UserRepo) DeleteUser(id string) error {
	var user = db_model.User{
		Id: uuid.MustParse(id),
	}

	res := u.db.Delete(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *UserRepo) PaginateUser(p *model.Pagination) ([]*db_model.User, error) {
	var users []*db_model.User
	res := u.db.Scopes(
		helper.Paginate(p, u.db),
	).Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}
