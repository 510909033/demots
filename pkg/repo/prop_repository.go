package repo

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type propRepo struct{}

func NewPropRepository() model.PropRepository {
	return &propRepo{}
}

func (r *propRepo) Create(ctx context.Context, tx *gorm.DB, prop *model.PropEntity) {
	vals, err := json.Marshal(&prop.Entity)
	fn.PanicErr(err)
	prop.Content = string(vals)

	err = tx.WithContext(ctx).Create(&prop).Error
	fn.PanicErr(err)
}

func (r *propRepo) Get(ctx context.Context, tx *gorm.DB, propId int64) model.Proper {
	var p model.PropEntity
	err := tx.WithContext(ctx).Where("id = ?", propId).First(&p).Error
	if err != nil {
		panic(err)
	}
	return &p
}
