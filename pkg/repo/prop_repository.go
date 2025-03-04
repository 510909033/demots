package repo

import (
	"api/pkg/model"
	"api/pkg/utils/fn"
	"context"
	"encoding/json"

	"gorm.io/gorm"
)

type propRepo struct{}

// GetRandomProp implements model.PropRepository.
func (r *propRepo) GetRandomProp(ctx context.Context, tx *gorm.DB) *model.PropEntity {
	var resp model.PropEntity
	err := tx.Model(&model.PropEntity{}).Order("RAND()").Find(&resp).Error
	fn.PanicErr(err)
	return &resp

}

func NewPropRepository() model.PropRepository {
	return &propRepo{}
}

func (r *propRepo) Create(ctx context.Context, tx *gorm.DB, prop *model.PropEntity) {
	vals, err := json.Marshal(prop.GetEntity())
	fn.PanicErr(err)
	prop.SetContent(string(vals))

	err = tx.WithContext(ctx).Create(prop).Error
	fn.PanicErr(err)
}

func (r *propRepo) Get(ctx context.Context, tx *gorm.DB, propId int64) *model.PropEntity {
	var p model.PropEntity
	err := tx.WithContext(ctx).Where("id = ?", propId).First(&p).Error
	if err != nil {
		panic(err)
	}
	return &p
}
