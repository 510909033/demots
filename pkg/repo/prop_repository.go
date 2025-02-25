package repo

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

type propRepo struct{}

func NewPropRepository() model.PropRepository {
	return &propRepo{}
}

func (r *propRepo) Create(ctx context.Context, tx *gorm.DB, prop *model.PropEntity) {
	err := tx.WithContext(ctx).Create(&prop).Error
	if err != nil {
		panic(err)
	}
}

func (r *propRepo) Get(ctx context.Context, tx *gorm.DB, propId int64) *model.PropEntity {
	var p model.PropEntity
	err := tx.WithContext(ctx).Where("id = ?", propId).First(&p).Error
	if err != nil {
		panic(err)
	}
	return &p
}
