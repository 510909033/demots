package model

import (
	"context"

	"gorm.io/gorm"
)

type PropEntity struct {
	// 获取道具id
	Id int64
	// 获取道具类型
	PropType int64
	// 道具单位
	// PropUnit int64
	// 有效期 开始时间戳，0表示不限
	StartTs int64
	// 有效期 结束时间戳，0表示不限
	EndTs int64
}

type PropRepository interface {
	Create(ctx context.Context, tx *gorm.DB, prop *PropEntity)
	Get(ctx context.Context, tx *gorm.DB, propId int64) *PropEntity
}
