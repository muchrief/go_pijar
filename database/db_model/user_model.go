package db_model

import (
	"github.com/google/uuid"
	"github.com/muchrief/go_pijar/src/helper"
	"gorm.io/gorm"
)

type UserRole string

const (
	ADMIN  UserRole = "admin"
	PUBLIC UserRole = "public"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `gorm:"not_null;unique" json:"username"`
	Password  string    `gorm:"not_null" json:"-"`
	Role      UserRole  `gorm:"default:public;not_null" json:"-"`
	CreatedAt int64     `gorm:"autoCreateTime:mili" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime:mili" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	u.Id = id

	u.Password = helper.Hash(u.Password)

	return nil
}
