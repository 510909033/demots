package model

import (
	"context"

	"gorm.io/gorm"
)

type UserBagRepository interface {
	Create(ctx context.Context, tx *gorm.DB, userBag *UserBagEntity)
	Delete(ctx context.Context, tx *gorm.DB, userId int64, propId int64)
	GetAll(ctx context.Context, tx *gorm.DB, userId int64) []*UserBagEntity
	// 更改数量
	SetCount(ctx context.Context, tx *gorm.DB, userId int64, propId int64, count int64)
}

type UserBagUseCase interface {
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
