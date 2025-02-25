package repo

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(d *gorm.DB) {
	db = d
}

func NewUser() model.UserRepository {
	return &User{}
}

type User struct {
	Id         int64  `json:"id" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"column:name;type:varchar(255);not null;default:''"`
	Experience int64  `json:"experience" gorm:"column:experience;type:bigint;not null;default:0"`
}

// Save implements model.User.
func (u *User) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.Create(u).Error
}

// AddExperience implements model.User.
func (u *User) AddExperience(ctx context.Context, experience int64) {
	// experience += u.Experience
	// return u.db.Model(u).Update("experience", gorm.Expr("experience + ?", experience)).Error
}

// GetId implements model.User.
func (u *User) GetId(ctx context.Context) int64 {
	// TODO W ?
	return u.Id
}

// SetName implements model.User.
func (u *User) SetName(ctx context.Context, name string) {
	u.Name = name
}
