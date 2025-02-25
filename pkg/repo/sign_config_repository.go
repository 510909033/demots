package repo

import (
	"api/pkg/model"
	"context"

	"gorm.io/gorm"
)

type signConfigRepo struct{}

func NewSignConfigRepository() model.SignConfigRepository {
	return &signConfigRepo{}
}

func (r *signConfigRepo) GetByScene(ctx context.Context, tx *gorm.DB, scene int64) *model.SignConfigEntity {
	var signConfig model.SignConfigEntity
	tx.Where("scene = ?", scene).First(&signConfig)
	return &signConfig
}

func (r *signConfigRepo) Create(ctx context.Context, tx *gorm.DB, signConfig *model.SignConfigEntity) {
	if err := tx.WithContext(ctx).Create(&signConfig).Error; err != nil {
		panic(err)
	}
}
