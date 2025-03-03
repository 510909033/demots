package repo

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"

	"gorm.io/gorm"
)

func NewGearRepository() model.GearRepository {
	return &gearRepo{}
}

type gearRepo struct{}

// Create implements model.GearRepository.
func (g *gearRepo) Create(ctx context.Context, tx *gorm.DB, param *model.GearEntity) {
	err := tx.WithContext(ctx).Create(&param).Error
	fn.PanicErr(err)
}
