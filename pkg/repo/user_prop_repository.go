package repo

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

type UserPropRepositoryImpl struct {
	db *gorm.DB
}

func NewUserPropRepository(db *gorm.DB) model.UserPropRepository {
	return &UserPropRepositoryImpl{db: db}
}

func (r *UserPropRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, userProp *model.UserPropEntity) {
	if err := tx.WithContext(ctx).Create(&userProp).Error; err != nil {
		panic(err)
	}
}
