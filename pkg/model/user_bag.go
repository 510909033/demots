package model

import (
	"context"

	"gorm.io/gorm"
)

type UserBager interface {
	GetUserId() int64
	GetPropId() int64
	GetAddTs() int64
	GetEndTs() int64
	GetCount() int64

	// set
	SetAddTs(addTs int64)
	SetEndTs(endTs int64)
	SetCount(count int64)
}

func NewUserBager(userid NewUserIder, propid NewPropId) UserBager {
	return (UserBager)(nil)
}

// // 用户背包
// type UserBag struct {
// 	UserID int64 `gorm:"primaryKey"`
// 	PropId int64 `gorm:"primaryKey"`

// 	// // 物品类型, 0未知， 1装备， 道具
// 	// BagType int64
// 	// // 物品id
// 	// ContentId int64
// 	// 添加时间戳
// 	AddTs int64
// 	// 有效期时间戳，0表示不失效
// 	EndTs int64
// 	// 物品数量
// 	Count int64
// }

type UserBagRepository interface {
	Create(ctx context.Context, tx *gorm.DB, userBag UserBager)
	Delete(ctx context.Context, tx *gorm.DB, userId int64, propId int64)
	GetAll(ctx context.Context, tx *gorm.DB, userId int64) []*UserBager
	// 更改数量
	SetCount(ctx context.Context, tx *gorm.DB, userId int64, propId int64, count int64)
}

type UserBagUseCase interface {
	// // 添加物品
	// AddUserBag(ctx context.Context, user *UserEntity, param *UserBag)
	// // 减少物品
	// SubUserBag(ctx context.Context, user *UserEntity, param *UserBag)
	// // 获取物品
	// GetUserBag(ctx context.Context, user *UserEntity, propId int64) *UserBag
	// // 获取所有物品
	// GetAllUserBag(ctx context.Context, user *UserEntity) []*UserBag
}

type UserBagEntity struct {
	UserId int64
	PropId int64
	AddTs  int64
	EndTs  int64
	Count  int64
}

func (u *UserBagEntity) GetUserId() int64 {
	return u.UserId
}

func (u *UserBagEntity) GetPropId() int64 {
	return u.PropId
}

func (u *UserBagEntity) GetAddTs() int64 {
	return u.AddTs
}

func (u *UserBagEntity) GetEndTs() int64 {
	return u.EndTs
}

func (u *UserBagEntity) GetCount() int64 {
	return u.Count
}

func (u *UserBagEntity) SetUserId(userId int64) {
	u.UserId = userId
}

func (u *UserBagEntity) SetPropId(propId int64) {
	u.PropId = propId
}

func (u *UserBagEntity) SetAddTs(addTs int64) {
	u.AddTs = addTs
}

func (u *UserBagEntity) SetEndTs(endTs int64) {
	u.EndTs = endTs
}

func (u *UserBagEntity) SetCount(count int64) {
	u.Count = count
}
