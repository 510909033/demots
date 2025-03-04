package repo

import (
	"context"

	"api/pkg/model"
	"api/pkg/utils/fn"

	"gorm.io/gorm"
)

type userBagRepo struct {
}

func NewUserBagRepository() model.UserBagRepository {
	return &userBagRepo{}
}

func (r *userBagRepo) Create(ctx context.Context, tx *gorm.DB, userBag *model.UserBagEntity) {
	err := tx.Create(&userBag).Error
	fn.PanicErr(err)
}

func (r *userBagRepo) Delete(ctx context.Context, tx *gorm.DB, userId int64, propId int64) {
	err := tx.Where("user_id = ? AND prop_id = ?", userId, propId).Delete(&model.UserBagEntity{}).Error
	fn.PanicErr(err)
}

func (r *userBagRepo) GetAll(ctx context.Context, tx *gorm.DB, userId int64) []*model.UserBagEntity {
	var userBags []*model.UserBagEntity
	err := tx.Where("user_id = ?", userId).Find(&userBags).Error
	fn.PanicErr(err)

	return userBags
}

func (r *userBagRepo) SetCount(ctx context.Context, tx *gorm.DB, userId int64, propId int64, count int64) {
	err := tx.Model(&model.UserBagEntity{}).Where("user_id = ? AND prop_id = ?", userId, propId).Update("count", count).Error
	fn.PanicErr(err)
}
